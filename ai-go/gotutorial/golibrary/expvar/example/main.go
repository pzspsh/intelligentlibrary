/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 16:58:44
*/
package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
)

func kvFunc(kv expvar.KeyValue) {
	fmt.Println(kv.Key, kv.Value)
}

/*
1）New*()  //  新建一个变量
2）Set(*)   //  设置这个变量
3）Add(*)  //  在原有变量上加上另一个变量
4）String()  // 实现Var接口
除此之外，Map还有几个特有的函数：
1）Init()                  // 初始化Map
2）Get(key string)  // 根据key获取value
3）Do(f func(Key Value))  // 对Map中的每对key/value执行函数f
*/
func main() {
	inerInt := int64(10)
	pubInt := expvar.NewInt("Int")
	pubInt.Set(inerInt)
	pubInt.Add(2)

	inerFloat := 1.2
	pubFloat := expvar.NewFloat("Float")
	pubFloat.Set(inerFloat)
	pubFloat.Add(0.1)

	inerString := "hello"
	pubString := expvar.NewString(inerString)
	pubString.Set(inerString)

	pubMap := expvar.NewMap("Map").Init()
	pubMap.Set("Int", pubInt)
	pubMap.Set("Float", pubFloat)
	pubMap.Set("String", pubString)
	pubMap.Do(kvFunc)
	pubMap.Add("Int", 1)
	pubMap.Add("NewInt", 123)
	pubMap.AddFloat("Float", 0.5)
	pubMap.AddFloat("NewFloat", 0.9)
	pubMap.Do(kvFunc)
	// pubMap.Delete("NewInt")

	expvar.Do(kvFunc)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
