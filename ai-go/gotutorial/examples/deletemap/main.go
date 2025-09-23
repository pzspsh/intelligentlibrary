package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	mu           sync.Mutex
	ApsAssetChan = make(chan string, 1000)
)

func StatsRun() error {
	var rwmap = make(map[string]bool, 1000)
	go func() {
		for {
			if rwnumber, ok := <-ApsAssetChan; ok {
				mu.Lock()
				if _, ok := rwmap[rwnumber]; !ok {
					rwmap[rwnumber] = true
				}
				mu.Unlock()
			}
		}
	}()
	for {
		if len(rwmap) > 0 {
			mu.Lock()
			for rwnumber := range rwmap {
				fmt.Println(rwnumber)
				delete(rwmap, rwnumber)
			}
			mu.Unlock()
			fmt.Println("====================================================")
			time.Sleep(100 * time.Microsecond)
		}
	}
}

func RandomStr(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func pushdata() {
	for {
		ApsAssetChan <- RandomStr(10)
		time.Sleep(100 * time.Microsecond)
	}
}

func main() {
	go pushdata()
	StatsRun()
}
