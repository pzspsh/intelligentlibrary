/*
@File   : connect.go
@Author : pan
@Time   : 2023-06-13 15:02:03
*/
package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func (r *RedisConfig) RedisConn() (*redis.Client, error) {
	opt := redis.Options{}
	opt = redis.Options{
		Addr:        fmt.Sprintf("%v:%v", r.Host, r.Port),
		Password:    r.Password,
		DB:          r.DB,
		DialTimeout: 5 * time.Second,
	}
	client := redis.NewClient(&opt)
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	} else {
		return client, nil
	}
	/*
					集群：
					func initClient()(err error){
					rdb := redis.NewClusterClient(&redis.ClusterOptions{
						Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
					})
					_, err = rdb.Ping().Result()
					if err != nil {
						return err
					}
					return nil
				}

				哨兵模式：
				func initClient()(err error){
			rdb := redis.NewFailoverClient(&redis.FailoverOptions{
				MasterName:    "master",
				SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
			})
			_, err = rdb.Ping().Result()
			if err != nil {
				return err
			}
			return nil
		}
	*/
}

func main() {
	r := &RedisConfig{
		Host:     "ip",
		Port:     "port",
		Password: "pass",
		DB:       0,
	}
	client, err := r.RedisConn()
	if err != nil {
		fmt.Printf("redis conn err:%v", err)
	} else {
		fmt.Printf("redis conn successful:%v", client)
	}
}
