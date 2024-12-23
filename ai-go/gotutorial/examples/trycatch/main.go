/*
@File   : main.go
@Author : pan
@Time   : 2023-11-14 10:08:25
*/
package main

import (
	"log"
)

type ExceptionStruct struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}
type Exception interface{}

func Throw(up Exception) {
	panic(up)
}
func (ex ExceptionStruct) Do() {
	if ex.Finally != nil {
		defer ex.Finally()
	}
	if ex.Catch != nil {
		defer func() {
			if e := recover(); e != nil {
				ex.Catch(e)
			}
		}()
	}
	ex.Try()
}

func main() {
	log.Println("开始执行...")
	ExceptionStruct{
		Try: func() {
			log.Println("try...")
			Throw("发生了错误")
		},
		Catch: func(e Exception) {
			log.Println("exception=", e)
		},
		Finally: func() {
			log.Println("Finally...")
		},
	}.Do()
	log.Println("结束运行")
}
