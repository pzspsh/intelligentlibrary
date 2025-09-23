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

func StatsRun() {
	var wg sync.WaitGroup
	var updatech = make(chan bool, 5)
	var taskchan = make(chan string, 10)
	var newlist []string
	for {
		var numberlist []string
		if len(newlist) > 0 {
			numberlist = newlist
			newlist = []string{}
		} else {
			numberlist = GenNumber()
		}
		if len(numberlist) > 0 {
			fmt.Println("开始生成任务", numberlist)
			for _, number := range numberlist {
				select {
				case taskchan <- number:
				default:
					newlist = append(newlist, number)
				}
			}
		} else {
			time.Sleep(time.Second * 10)
		}
		for range len(taskchan) {
			wg.Add(1)
			number := <-taskchan
			updatech <- true
			go func(number string, update chan bool) {
				defer wg.Done()
				defer func() { <-update }()
				fmt.Println(number)
			}(number, updatech)
		}
		wg.Wait()
	}
}

func StatsRun2() {
	var wg sync.WaitGroup
	var m sync.Map
	var updatech = make(chan bool, 5)
	var taskchan = make(chan string, 10)
	go func() {
		var newlist []string
		for {
			var numberlist []string
			if len(newlist) > 0 {
				numberlist = newlist
				newlist = []string{}
			} else {
				numberlist = GenNumber()
			}
			if len(numberlist) > 0 {
				fmt.Println("开始生成任务", numberlist)
				for _, number := range numberlist {
					select {
					case taskchan <- number:
						m.Store(number, true)
					default:
						newlist = append(newlist, number)
					}
				}
			} else {
				time.Sleep(time.Second * 10)
			}
		}
	}()

	for {
		if len(taskchan) > 0 {
			for number := range taskchan {
				wg.Add(1)
				updatech <- true
				go func(number string, update chan bool) {
					defer wg.Done()
					defer func() { <-update }()
					fmt.Println(number)
					m.Delete(number)
				}(number, updatech)
			}
			wg.Wait()
		} else {
			time.Sleep(5 * time.Second)
		}

	}
}

func StatsRun3() {
	var updatech = make(chan bool, 5)
	var taskchan = make(chan string, 10)
	go func() {
		for {
			numberlist := GenNumber()
			if len(numberlist) > 0 {
				fmt.Println("开始生成任务", numberlist)
				for _, number := range numberlist {
					taskchan <- number
				}
			} else {
				time.Sleep(time.Second * 10)
			}
			time.Sleep(10 * time.Second)
		}
	}()

	for {
		if number, ok := <-taskchan; ok {
			updatech <- true
			go func(number string, update chan bool) {
				defer func() { <-update }()
				fmt.Println(number)
				time.Sleep(2 * time.Second)
			}(number, updatech)
		}
	}
}

func main() {
	// StatsRun()
	StatsRun2()
	// StatsRun3()
}
