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
