/*
@File   : main.go
@Author : pan
@Time   : 2025-06-13 16:04:59
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func RandomStr(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenNumber() []string {
	var number []string
	for range 20 {
		number = append(number, RandomStr(10))
	}
	return number
}

func main() {
	var m sync.Map
	var wg sync.WaitGroup
	updatech := make(chan int, 2)
	var countint = 0
	go func() {
		for {
			var ranstr string
			var numberlist = []string{"number1", "number2", "number3", "number4"}
			if countint%2 == 0 {
				ranstr = RandomStr(10)
				numberlist = append(numberlist, ranstr)
			}
			for _, number := range numberlist {
				if _, ok := m.Load(number); !ok {
					m.Store(number, 1)
				}
			}
			countint++
			time.Sleep(time.Second * 2)
		}
	}()
	for {
		var hasValue bool
		m.Range(func(key, value any) bool {
			hasValue = true
			return false
		})
		if hasValue {
			m.Range(func(key, value any) bool { // 遍历所有键值对
				wg.Add(1)
				updatech <- 1
				go func(update chan int) {
					defer wg.Done()
					defer func() { <-update }()
					m.Delete(key)
					fmt.Println(key.(string), value.(int))
					time.Sleep(time.Second * 5)
				}(updatech)
				hasValue = true
				return true
			})
			wg.Wait()
			fmt.Println("============================================================================")
		} else {
			time.Sleep(time.Second * 2)
		}
	}
}
