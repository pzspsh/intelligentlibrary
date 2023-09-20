/*
@File   : tokenbucket.go
@Author : pan
@Time   : 2023-09-20 15:15:33
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// 并发访问同一个user_id/ip的记录需要上锁
var recordMu map[string]*sync.RWMutex

func init() {
	recordMu = make(map[string]*sync.RWMutex)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TokenBucket struct {
	BucketSize int                // 木桶内的容量：最多可以存放多少个令牌
	TokenRate  time.Duration      // 多长时间生成一个令牌
	records    map[string]*record // 报错user_id/ip的访问记录
}

// 上次访问时的时间戳和令牌数
type record struct {
	last  time.Time
	token int
}

func NewTokenBucket(bucketSize int, tokenRate time.Duration) *TokenBucket {
	return &TokenBucket{
		BucketSize: bucketSize,
		TokenRate:  tokenRate,
		records:    make(map[string]*record),
	}
}

func (t *TokenBucket) getUidOrIp() string {
	// 获取请求用户的user_id或者ip地址
	return "127.0.0.1"
}

// 获取这个user_id/ip上次访问时的时间戳和令牌数
func (t *TokenBucket) getRecord(uidOrIp string) *record {
	if r, ok := t.records[uidOrIp]; ok {
		return r
	}
	return &record{}
}

// 保存user_id/ip最近一次请求时的时间戳和令牌数量
func (t *TokenBucket) storeRecord(uidOrIp string, r *record) {
	t.records[uidOrIp] = r
}

// 验证是否能获取一个令牌
func (t *TokenBucket) validate(uidOrIp string) bool {
	// 并发修改同一个用户的记录上写锁
	rl, ok := recordMu[uidOrIp]
	if !ok {
		var mu sync.RWMutex
		rl = &mu
		recordMu[uidOrIp] = rl
	}
	rl.Lock()
	defer rl.Unlock()

	r := t.getRecord(uidOrIp)
	now := time.Now()
	if r.last.IsZero() {
		// 第一次访问初始化为最大令牌数
		r.last, r.token = now, t.BucketSize
	} else {
		if r.last.Add(t.TokenRate).Before(now) {
			// 如果与上次请求的间隔超过了token rate
			// 则增加令牌，更新last
			r.token += max(int(now.Sub(r.last)/t.TokenRate), t.BucketSize)
			r.last = now
		}
	}
	var result bool
	if r.token > 0 {
		// 如果令牌数大于1，取走一个令牌，validate结果为true
		r.token--
		result = true
	}

	// 保存最新的record
	t.storeRecord(uidOrIp, r)
	return result
}

// 返回是否被限流
func (t *TokenBucket) IsLimited() bool {
	return !t.validate(t.getUidOrIp())
}

func main() {
	tokenBucket := NewTokenBucket(5, 100*time.Millisecond)
	for i := 0; i < 6; i++ {
		fmt.Println(tokenBucket.IsLimited())
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println(tokenBucket.IsLimited())
}
