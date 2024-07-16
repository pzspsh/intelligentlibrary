/*
@File   : main.go
@Author : pan
@Time   : 2024-07-16 11:05:26
*/
package main

/*
"低耦合"是软件设计中另一个重要的原则，它指的是模块之间应该尽可能独立，一个模块的功能变化或修改应该尽可能少地
影响其他模块。低耦合的系统通常更加健壮、可维护，并且更易于扩展。

在Go语言中，实现低耦合的代码通常意味着遵循以下原则：

明确的接口：模块之间通过明确的接口进行通信，而不是直接依赖于对方的具体实现。
松散的依赖：模块之间的依赖关系应该尽可能少，且只依赖必要的部分。
模块自治：每个模块应该能够独立地工作，而不依赖于其他模块的状态或行为。
下面是一个使用Go语言实现低耦合的例子，这个例子展示了一个简单的日志记录系统，其中包含了日志记录器和日志处理器两个模块。
*/
import (
	"algorithm/often/lowcoupling/logger"
	"algorithm/often/lowcoupling/logprocessor"
)

func main() {
	// 创建一个ConsoleLogger实例
	consoleLogger := &logger.ConsoleLogger{}

	// 使用ConsoleLogger创建一个LogProcessor实例
	logProcessor := logprocessor.NewLogProcessor(consoleLogger)

	// 处理一条日志消息
	logProcessor.ProcessLog("This is a log message.")
}
