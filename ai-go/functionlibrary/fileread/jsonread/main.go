/*
@File   : main.go
@Author : pan
@Time   : 2024-03-01 17:34:03
*/
package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    data := `{"name":"John","age":31}
{"name":"John","age":31}`

    // 使用 "\n" 分割字符串以获得每个 JSON 对象
    objects := strings.Split(data, "\n")

    var people []Person

    for _, obj := range objects {
        if obj == "" {
            continue
        }
        var person Person
        err := json.Unmarshal([]byte(obj), &person)
        if err != nil {
            fmt.Println("解析错误:", err)
            continue
        }
        people = append(people, person)
    }

    fmt.Println(people)
}