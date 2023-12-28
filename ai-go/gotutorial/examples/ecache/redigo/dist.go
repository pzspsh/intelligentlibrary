/*
@File   : dist.go
@Author : pan
@Time   : 2023-12-28 11:25:16
*/
package dist

import (
	"log"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"gotutorial/examples/ecache"
)

const topic = "orca-zhang/ecache"

// `RedisCli`` interface used by `dist` component
type RedisCli interface {
	// if the redis client is ready
	OK() bool
	// pub a payload to channel
	Pub(channel, payload string) error
	// sub a payload from channel, callback uill tidy the local cache
	Sub(channel string, callback func(payload string)) error
}

var redisCli RedisCli
var m sync.Map

func delAll(pool, key string) {
	if caches, _ := m.Load(pool); caches != nil {
		for _, c := range *(caches.(*[]*ecache.Cache)) {
			c.Del(key)
		}
	}
}

// Init `dist` component with redis client
func Init(r RedisCli) {
	if redisCli != r {
		redisCli = r
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Println(err)
					debug.PrintStack()
				}
			}()

			for {
				for r == nil || !r.OK() {
					time.Sleep(10 * time.Millisecond)
				}
				_ = r.Sub(topic, func(payload string) {
					vs := strings.Split(payload, ":")
					if len(vs) >= 2 {
						delAll(vs[0], vs[1])
					}
				})
			}
		}()
	}
}

// Bind - to enable distributed consistency
// `pool` is not necessary, it can be used to classify instances that store same items, but it will be more efficient if it is not empty
// `caches` is cache instances to be binded
func Bind(pool string, caches ...*ecache.Cache) error {
	c, _ := m.LoadOrStore(pool, &[]*ecache.Cache{})
	*(c.(*[]*ecache.Cache)) = append(*(c.(*[]*ecache.Cache)), caches...)
	return nil
}

// OnDel - delete `key` in `pool` at distributed scale
func OnDel(pool, key string) error {
	// pub to remote nodes
	r := redisCli
	if r != nil && r.Pub(topic, strings.Join([]string{pool, key}, ":")) == nil {
		return nil
	}
	delAll(pool, key)
	return nil
}