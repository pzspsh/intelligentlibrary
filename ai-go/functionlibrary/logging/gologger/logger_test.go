/*
@File   : logger_test.go
@Author : pan
@Time   : 2023-10-25 16:17:25
*/
package logger_test

import (
	logger "function/logging/gologger"
	"testing"
)

func TestLogger(t *testing.T) {
	logger.Logger(`..\data\log\log.txt`, 10, "mb")
	log1 := logger.LoggerSet(`..\data\log\log1.txt`, 10, "10mb")
	log2 := logger.LoggerSet(`..\data\log\log2.txt`, 10, "10mb")
	logger.Info("logger hello world")
	logger.Info("log hello world")
	log1.Error("log1 你好！！！！！！！")
	log2.Error("log2 你好！！！！！！！")
}
