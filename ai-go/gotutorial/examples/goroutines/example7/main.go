package main

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"runtime"
	"sync"
)

type StationStats struct {
	min, max, sum, count int
}

func (s *StationStats) Update(temp int) {
	if s.count == 0 {
		s.min, s.max = temp, temp
	}
	if temp < s.min {
		s.min = temp
	}
	if temp > s.max {
		s.max = temp
	}
	s.sum += temp
	s.count++
}

// ---- 自定义哈希表 ----
type entry struct {
	key   []byte
	stats StationStats
	used  bool
}

type hashTable struct {
	buckets []entry
}

func newHashTable(size int) *hashTable {
	return &hashTable{buckets: make([]entry, size)}
}

func fnv1a(data []byte) uint64 {
	h := fnv.New64a()
	h.Write(data)
	return h.Sum64()
}

func (h *hashTable) getOrInsert(key []byte) *StationStats {
	hash := fnv1a(key)
	idx := int(hash & uint64(len(h.buckets)-1))
	for {
		e := &h.buckets[idx]
		if !e.used {
			e.key = append([]byte(nil), key...)
			e.used = true
			return &e.stats
		}
		if bytes.Equal(e.key, key) {
			return &e.stats
		}
		idx = (idx + 1) & (len(h.buckets) - 1)
	}
}

// ---- 高性能温度解析 ----

func parseTemp(b []byte) int {
	neg := false
	if b[0] == '-' {
		neg = true
		b = b[1:]
	}
	index := 0
	temp := int(b[0] - '0')
	index++
	if b[index] != '.' {
		temp = temp*10 + int(b[index]-'0')
		index++
	}
	index++
	temp = temp*10 + int(b[index]-'0')
	if neg {
		return -temp
	}
	return temp
}

func processChunk(filename string, start, end int64) *hashTable {
	f, _ := os.Open(filename)
	defer f.Close()
	// 确保从行首开始（除第一个块外）
	if start > 0 {
		f.Seek(start, io.SeekStart)
		// 读取直到下一个换行符
		buf := make([]byte, 1)
		for {
			n, err := f.Read(buf)
			if n == 0 || err != nil || buf[0] == '\n' {
				break
			}
			start++
		}
		// 此时start指向新行的开始位置
	} else {
		f.Seek(start, io.SeekStart)
	}
	buf := make([]byte, 1<<20)
	// 1MB
	ht := newHashTable(1 << 20)
	leftover := []byte{}
	for {
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		data := append(leftover, buf[:n]...)
		lines := bytes.Split(data, []byte("\n"))
		if err != io.EOF {
			leftover = lines[len(lines)-1]
			lines = lines[:len(lines)-1]
		}
		for _, line := range lines {
			if len(line) == 0 {
				continue
			}
			idx := bytes.LastIndexByte(line, ';')
			station := line[:idx]
			temp := parseTemp(line[idx+1:])
			s := ht.getOrInsert(station)
			s.Update(temp)
		}
		if err == io.EOF {
			break
		}
	}
	return ht
}
func optimized(filename string, workers int) (map[string]*StationStats, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fi, _ := f.Stat()
	size := fi.Size()
	chunk := size / int64(workers)
	results := make([]*hashTable, workers)
	wg := sync.WaitGroup{}
	for i := range workers {
		start := int64(i) * chunk
		end := start + chunk
		if i == workers-1 {
			end = size
		}
		wg.Add(1)
		go func(idx int, s, e int64) {
			defer wg.Done()
			results[idx] = processChunk(filename, s, e)
		}(i, start, end)
	}
	wg.Wait()
	// 合并结果
	final := make(map[string]*StationStats)
	for _, ht := range results {
		for _, e := range ht.buckets {
			if !e.used {
				continue
			}
			k := string(e.key)
			s, ok := final[k]
			if !ok {
				s = &StationStats{}
				final[k] = s
			}
			s.sum += e.stats.sum
			s.count += e.stats.count
			if e.stats.min < s.min || s.count == e.stats.count {
				s.min = e.stats.min
			}
			if e.stats.max > s.max || s.count == e.stats.count {
				s.max = e.stats.max
			}
		}
	}
	return final, nil
}
func main() {
	res, _ := optimized("measurements.txt", runtime.NumCPU())
	for k, v := range res {
		avg := float64(v.sum) / float64(v.count) / 10.0
		fmt.Printf("%s=%.1f (min: %.1f, max: %.1f)\n", k, avg, float64(v.min)/10, float64(v.max)/10)
	}
}
