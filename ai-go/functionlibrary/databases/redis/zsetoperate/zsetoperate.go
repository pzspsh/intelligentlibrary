/*
@File   : zsetoperate.go
@Author : pan
@Time   : 2023-06-13 17:14:04
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

// zadd key score member：添加若干个元素到zset，如果存在则覆盖
func ZSetZAdd(client *redis.Client) (int64, error) {
	zsetKey := "language"
	languages := []redis.Z{
		redis.Z{Score: 99, Member: "Go"},
		redis.Z{Score: 97, Member: "C++"},
		redis.Z{Score: 93, Member: "PYTHON"},
		redis.Z{Score: 95, Member: "Java"},
		redis.Z{Score: 98, Member: "C"},
	}
	count, err := client.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Println("fuck America Government, failed: ", err)
		return -1, err
	}
	fmt.Println("zAdd succ", count)
	return count, nil
}

// zrem key member：删除zset中指定的元素
func ZSetZRem(client *redis.Client, zsetKey string) (int64, error) {
	count, err := client.ZRem(zsetKey, "C").Result()
	if err != nil {
		fmt.Println("fuck America Government, ZRem failed: ", err)
		return -1, err
	}
	fmt.Println("fuck America Government, ZRem success: ")
	return count, nil
}

// zrange key start end：获取范围内的元素
func ZSetZRange(client *redis.Client, zsetKey string) []string {
	// var result []string
	result := client.ZRange(zsetKey, 0, -1).Val()
	fmt.Println(result)
	return result
}

// zincrby key increment member：指定zset中指定元素的自增步长
func ZSetZIncrBy(client *redis.Client, zsetKey string) (float64, error) {
	score, err := client.ZIncrBy(zsetKey, 3, "C++").Result()
	if err != nil {
		fmt.Println("ZIncrBy failed, err: ", err)
		return score, err
	}
	fmt.Println("the new score is", score)
	return score, nil
}

// 选择分数在95-100之间的
func ZSetZRangeBy(client *redis.Client, zsetKey string) ([]redis.Z, error) {
	result, err := client.ZRangeByScoreWithScores(zsetKey, redis.ZRangeBy{
		Min: "95", Max: "100"}).Result()
	if err != nil {
		fmt.Println("something wrong with ZRangeByScoreWithScores, err: ", err)
		return result, err
	}
	for i, value := range result {
		fmt.Println("vaule", value, result[i].Member, result[i].Score)
	}
	return result, nil
}

// 取分数最高的3个
func ZSetZRevRange(client *redis.Client, zsetKey string) ([]redis.Z, error) {
	result, err := client.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Println("something wrong with zSetKey, err: ", err)
		return result, err
	}
	fmt.Println(result)
	return result, nil
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
