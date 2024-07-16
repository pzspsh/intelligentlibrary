/*
@File   : logprocessor.go
@Author : pan
@Time   : 2024-07-16 11:07:25
*/
package logprocessor

import (
	"algorithm/often/lowcoupling/logger" // 导入logger包，但不直接使用ConsoleLogger

	"fmt"
)

/*
接下来，我们定义一个LogProcessor，它负责处理日志消息。注意，LogProcessor不直接依赖于ConsoleLogger或其他具体的日
志记录器，而是依赖于Logger接口：
*/

// LogProcessor 处理日志消息
type LogProcessor struct {
	logger logger.Logger // 依赖于Logger接口，而不是具体的实现
}

// NewLogProcessor 创建一个新的LogProcessor实例，并注入一个Logger
func NewLogProcessor(l logger.Logger) *LogProcessor {
	return &LogProcessor{logger: l}
}

// ProcessLog 处理一条日志消息
func (p *LogProcessor) ProcessLog(message string) {
	// 在这里可以添加额外的处理逻辑，比如过滤、格式化等
	err := p.logger.Log(message) // 调用Logger接口的Log方法记录日志
	if err != nil {
		fmt.Println("Error logging message:", err)
	}
}
