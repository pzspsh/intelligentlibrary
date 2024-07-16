/*
@File   : logger.go
@Author : pan
@Time   : 2024-07-16 11:06:32
*/
package logger

import (
	"fmt"
)

// 首先，我们定义一个Logger接口，用于规范日志记录器的行为：
// Logger 定义了一个日志记录器的接口
type Logger interface {
	Log(message string) error
}

// 然后，我们实现一个具体的日志记录器ConsoleLogger，它实现了Logger接口，并将日志输出到控制台：
// ConsoleLogger 是一个简单的控制台日志记录器
type ConsoleLogger struct{}

// Log 实现了Logger接口的Log方法，将日志输出到控制台
func (l *ConsoleLogger) Log(message string) error {
	fmt.Println("Logging to console:", message)
	return nil
}
