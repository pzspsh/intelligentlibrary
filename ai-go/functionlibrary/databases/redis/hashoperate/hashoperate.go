/*
@File   : hashoperate.go
@Author : pan
@Time   : 2023-06-13 17:01:07
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

// hset key field value：添加或修改hash类型key的field的值
func HSetOperate(client *redis.Client) error {
	err := client.HSet("USA", "name", "dsb").Err()
	if err != nil {
		return err
	} else {
		return nil
	}
}

// hget key field：获取一个hash类型的key的field的值
func HGetOperate(client *redis.Client) (string, error) {
	result, err := client.HGet("USA", "name").Result()
	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}

// hmset key field1 value1 field2 value2：批量添加多个hash类型key的field值
func MapMSet(client *redis.Client) error {
	usaMap := map[string]interface{}{"name": "dsb", "name2": "robber"}
	err := client.HMSet("USA", usaMap).Err()
	if err != nil {
		return err
	} else {
		return nil
	}
}

// hmget key field1 field2：批量获取hash类型key多个field的value值
func MGet(client *redis.Client) []interface{} {
	result := client.HMGet("USA", "name2", "name").Val()
	return result
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
