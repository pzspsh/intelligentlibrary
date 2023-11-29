/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 17:09:39
*/
package main

import (
	"context"
	"fmt"
)

type Config struct {
	LogLevel string
	Timeout  int
}

// 模拟获取系统配置信息的函数
func getConfig(ctx context.Context) Config {
	// 从 ctx 中获取配置信息
	config, ok := ctx.Value("config").(Config)
	if !ok {
		return Config{LogLevel: "debug", Timeout: 100}
	}
	return config
}

func main() {
	// 初始化一个 context
	ctx := context.Background()
	// 设置系统配置信息到 context 中
	config := Config{LogLevel: "info", Timeout: 200}
	ctx = context.WithValue(ctx, "config", config)
	// 测试获取配置信息
	c := getConfig(ctx)
	fmt.Println("LogLevel:", c.LogLevel)
	fmt.Println("Timeout:", c.Timeout)
}
