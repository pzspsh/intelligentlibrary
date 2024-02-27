/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 09:58:49
*/
package main

import (
	"fmt"
	"hash/crc32"
	"log"
	"sort"
	"strconv"
	"strings"
)

// HashFunc 定义生成哈希的函数
type HashFunc func(data []byte) uint32

// Map 存储节点，你可以从Map中选择节点
type Map struct {
	hashFunc HashFunc       //哈希算法
	replicas int            //虚拟节点的数量
	keys     []int          //所有虚拟节点的哈希值
	hashMap  map[int]string //key->虚拟节点的哈希值，value->节点，实现虚拟节点映射到真实节点
}

// New 创建Map
func New(replicas int, fn HashFunc) *Map {
	m := &Map{
		replicas: replicas,
		hashFunc: fn,
		hashMap:  make(map[int]string),
	}
	if m.hashFunc == nil {
		m.hashFunc = crc32.ChecksumIEEE
	}
	return m
}

// IsEmpty Map中是否存在节点
func (m *Map) IsEmpty() bool {
	return len(m.keys) == 0
}

// AddNode 将给定的节点添加到一致性哈希中
func (m *Map) AddNode(keys ...string) {
	for _, key := range keys {
		if key == "" {
			continue
		}
		for i := 0; i < m.replicas; i++ {
			//计算虚拟节点哈希值
			hash := int(m.hashFunc([]byte(strconv.Itoa(i) + key)))
			//存储虚拟节点的哈希值
			m.keys = append(m.keys, hash)
			//存入map做映射
			m.hashMap[hash] = key
		}
	}
	//排序哈希值，下面匹配的时候要二分搜索
	sort.Ints(m.keys)
}

// getPartitionKey 支持哈希标记
func getPartitionKey(key string) string {
	beg := strings.Index(key, "{")
	if beg == -1 {
		return key
	}
	end := strings.Index(key, "}")
	if end == -1 || end == beg+1 {
		return key
	}
	return key[beg+1 : end]
}

// PickNode 获取与key最接近的节点。
func (m *Map) PickNode(key string) string {
	if m.IsEmpty() {
		return ""
	}
	partitionKey := getPartitionKey(key)
	//计算传入key的哈希
	hash := int(m.hashFunc([]byte(partitionKey)))
	// sort.Search 使用二分查找满足 m.keys[i] >= hash 的最小哈希值
	idx := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hash })
	// 若 key 的 hash 值大于最后一个虚拟节点的 hash 值，则选择第一个虚拟节点
	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]
}

func TestHash() {
	m := New(3, nil)
	m.AddNode("a", "b", "c", "d")
	if m.PickNode("zxc") != "a" {
		log.Println("wrong answer")
	}
	if m.PickNode("123{abc}") != "b" {
		log.Println("wrong answer")
	}
	if m.PickNode("abc") != "b" {
		log.Println("wrong answer")
	}
	for i := 0; i < 26; i++ {
		fmt.Println(fmt.Sprint(97+i)+"wxfQQ68725032", "在[", m.PickNode(fmt.Sprint(97+i)), "]节点")
	}
}

func main() {
	TestHash()
}
