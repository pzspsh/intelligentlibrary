/*
@File   : main.go
@Author : pan
@Time   : 2024-10-31 10:22:20
*/
package main

import (
	"flag"
	"fmt"
)

type Options struct {
	Str1   string
	Array1 strSlice
	Array2 strSlice
	Bool1  bool
}

type strSlice []string

func (s *strSlice) String() string {
	return fmt.Sprintf("%v", *s)
}

func (i *strSlice) Set(value string) error {
	var err error
	*i = append(*i, value)
	return err
}

var (
	options = &Options{}
)

func getParam() {
	flag.StringVar(&options.Str1, "s", "", "string的类型参数，默认为空“”")
	flag.Var(&options.Array1, "a1", "数组类型的参数，默认为[]")
	flag.Var(&options.Array2, "a2", "数组类型的参数，默认为[]")
	flag.BoolVar(&options.Bool1, "b", false, "bool类型的参数，默认为false")
	flag.Parse()
}

func main() {
	getParam()
	fmt.Println(options)
}

// go run main.go -s "hello world" -a1 "你" -a1 "好" -b  //结果：&{hello world [你 好] [] true}
