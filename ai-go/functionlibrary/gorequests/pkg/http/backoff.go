/*
@File   : backoff.go
@Author : pan
@Time   : 2023-08-24 15:38:44
*/
package http

import (
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Backoff func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration

func DefaultBackoff() func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	return func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
		mult := math.Pow(2, float64(attemptNum)) * float64(min)

		sleep := time.Duration(mult)
		if float64(sleep) != mult || sleep > max {
			sleep = max
		}
		return sleep
	}
}

func LinearJitterBackoff() func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	randMutex := &sync.Mutex{}

	return func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
		// attemptNum always starts at zero but we want to start at 1 for multiplication
		attemptNum++

		if max <= min {
			// Unclear what to do here, or they are the same, so return min *
			// attemptNum
			return min * time.Duration(attemptNum)
		}
		randMutex.Lock()
		jitter := rand.Float64() * float64(max-min)
		randMutex.Unlock()

		jitterMin := int64(jitter) + int64(min)
		return time.Duration(jitterMin * int64(attemptNum))
	}
}

func FullJitterBackoff() func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	randMutex := &sync.Mutex{}

	return func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
		duration := attemptNum * 1000000000 << 1

		randMutex.Lock()
		jitter := rand.Intn(duration-attemptNum) + int(min)
		randMutex.Unlock()

		if jitter > int(max) {
			return max
		}

		return time.Duration(jitter)
	}
}

func ExponentialJitterBackoff() func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	randMutex := &sync.Mutex{}

	return func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
		minf := float64(min)
		mult := math.Pow(2, float64(attemptNum)) * minf

		randMutex.Lock()
		jitter := rand.Float64() * (mult - minf)
		randMutex.Unlock()

		mult = mult + jitter

		sleep := time.Duration(mult)
		if sleep > max {
			sleep = max
		}
		return sleep
	}
}
