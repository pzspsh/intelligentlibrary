/*
@File   : setget.go
@Author : pan
@Time   : 2023-06-13 16:31:38
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

// set操作
func SetOperate(client *redis.Client) error {
	err := client.Set("golang", "i like", 0).Err
	if err() != nil {
		return err()
	} else {
		return nil
	}
}

// 批量set操作
func MSetOperate(client *redis.Client) error {
	err := client.MSet("golang", "value1", "c", "value2", "java", "value3", "python", "value4").Err
	if err() != nil {
		return err()
	} else {
		return nil
	}
}

// Get操作
func GetOperate(client *redis.Client) (string, error) {
	result, err := client.Get("golang").Result()
	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}

// 批量Get操作
func MGetOperate(client *redis.Client) ([]interface{}, error) {
	var result []interface{}
	var err error
	result, err = client.MGet("golang", "c", "java", "python").Result()
	for i, value := range result {
		fmt.Println(i, value)
	}
	if err != nil {
		return result, err
	} else {
		return result, nil
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
