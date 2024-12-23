/*
@File   : listoperate.go
@Author : pan
@Time   : 2023-06-13 16:51:08
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

// 从左边添加元素
func LPushOperate(client *redis.Client) error {
	err := client.LPush("足球", "贝利", "pan").Err
	if err() != nil {
		return err()
	} else {
		return nil
	}
}

// 获取list中的所有元素
func GetListdata(client *redis.Client) []string {
	val := client.LRange("足球", 0, -1).Val()
	return val
}

// 从list的左边弹出一个元素
func LpopOperate(client *redis.Client) (string, error) {
	result, err := client.LPop("足球").Result()
	if err != nil {
		return "", err
	} else {
		return result, err
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
