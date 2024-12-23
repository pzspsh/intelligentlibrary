/*
@File   : setoperate.go
@Author : pan
@Time   : 2023-06-13 17:13:34
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

// 向set中添加若干个元素
func SetSAdd(client *redis.Client) error {
	err := client.SAdd("Chinese", "AiGuo", "JingYe", "ChengXin", "YouShan").Err()
	if err != nil {
		fmt.Println("Are you right? Chinese people are really great", err)
		return err
	}
	return nil
}

// 返回set中所有元素s
func SetSmembers(client *redis.Client) []string {
	result := client.SMembers("Chinese").Val()
	fmt.Println(result)
	return result
}

// 判断元素是否存在于set中
func SetSismeber(client *redis.Client) (bool, error) {
	ok, err := client.SIsMember("Chinese", "AiGuo").Result() // 这里返回的是bool
	if err != nil {
		fmt.Println("You're mistaken")
		return ok, err
	}
	fmt.Println("I can see that the Chinese are really AiGuo is", ok)
	return ok, nil
}

// 返回set中元素的个数
func SetSCard(client *redis.Client) int64 {
	count := client.SCard("Chinese").Val()
	fmt.Printf("The Chinese people have %d virtues\n", count)
	return count
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
