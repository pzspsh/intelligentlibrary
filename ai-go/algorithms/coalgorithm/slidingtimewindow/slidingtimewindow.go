/*
@File   : slidingtimewindow.go
@Author : pan
@Time   : 2023-09-20 15:18:08
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var winMu map[string]*sync.RWMutex

func init() {
	winMu = make(map[string]*sync.RWMutex)
}

type timeSlot struct {
	timestamp time.Time // 这个timeSlot的时间起点
	count     int       // 落在这个timeSlot内的请求数
}

func countReq(win []*timeSlot) int {
	var count int
	for _, ts := range win {
		count += ts.count
	}
	return count
}

type SlidingWindowLimiter struct {
	SlotDuration time.Duration // time slot的长度
	WinDuration  time.Duration // sliding window的长度
	numSlots     int           // window内最多有多少个slot
	windows      map[string][]*timeSlot
	maxReq       int // win duration内允许的最大请求数
}

func NewSliding(slotDuration time.Duration, winDuration time.Duration, maxReq int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		SlotDuration: slotDuration,
		WinDuration:  winDuration,
		numSlots:     int(winDuration / slotDuration),
		windows:      make(map[string][]*timeSlot),
		maxReq:       maxReq,
	}
}

// 获取user_id/ip的时间窗口
func (l *SlidingWindowLimiter) getWindow(uidOrIp string) []*timeSlot {
	win, ok := l.windows[uidOrIp]
	if !ok {
		win = make([]*timeSlot, 0, l.numSlots)
	}
	return win
}

func (l *SlidingWindowLimiter) storeWindow(uidOrIp string, win []*timeSlot) {
	l.windows[uidOrIp] = win
}

func (l *SlidingWindowLimiter) validate(uidOrIp string) bool {
	// 同一user_id/ip并发安全
	mu, ok := winMu[uidOrIp]
	if !ok {
		var m sync.RWMutex
		mu = &m
		winMu[uidOrIp] = mu
	}
	mu.Lock()
	defer mu.Unlock()

	win := l.getWindow(uidOrIp)
	now := time.Now()
	// 已经过期的time slot移出时间窗
	timeoutOffset := -1
	for i, ts := range win {
		if ts.timestamp.Add(l.WinDuration).After(now) {
			break
		}
		timeoutOffset = i
	}
	if timeoutOffset > -1 {
		win = win[timeoutOffset+1:]
	}

	// 判断请求是否超限
	var result bool
	if countReq(win) < l.maxReq {
		result = true
	}

	// 记录这次的请求数
	var lastSlot *timeSlot
	if len(win) > 0 {
		lastSlot = win[len(win)-1]
		if lastSlot.timestamp.Add(l.SlotDuration).Before(now) {
			lastSlot = &timeSlot{timestamp: now, count: 1}
			win = append(win, lastSlot)
		} else {
			lastSlot.count++
		}
	} else {
		lastSlot = &timeSlot{timestamp: now, count: 1}
		win = append(win, lastSlot)
	}

	l.storeWindow(uidOrIp, win)

	return result
}

func (l *SlidingWindowLimiter) getUidOrIp() string {
	return "127.0.0.1"
}

func (l *SlidingWindowLimiter) IsLimited() bool {
	return !l.validate(l.getUidOrIp())
}

func main() {
	limiter := NewSliding(100*time.Millisecond, time.Second, 10)
	for i := 0; i < 5; i++ {
		fmt.Println(limiter.IsLimited())
	}
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 5; i++ {
		fmt.Println(limiter.IsLimited())
	}
	fmt.Println(limiter.IsLimited())
	for _, v := range limiter.windows[limiter.getUidOrIp()] {
		fmt.Println(v.timestamp, v.count)
	}

	fmt.Println("a thousand years later...")
	time.Sleep(time.Second)
	for i := 0; i < 7; i++ {
		fmt.Println(limiter.IsLimited())
	}
	for _, v := range limiter.windows[limiter.getUidOrIp()] {
		fmt.Println(v.timestamp, v.count)
	}
}
