/*
@File   : pipeline.go
@Author : pan
@Time   : 2023-06-13 17:14:35
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

func PipelineOperate(client *redis.Client) {
	pipeline := client.Pipeline()
	pipeline.Set("key1", "val1", time.Hour)
	pipeline.Set("key2", "val2", time.Hour)
	pipeline.Set("key3", "val3", time.Hour)
	pipeline.Set("key4", "val4", time.Hour)
	pipeline.Set("key5", "val5", time.Hour)
	_, err := pipeline.Exec()
	if err != nil {
		fmt.Println("set success")
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
