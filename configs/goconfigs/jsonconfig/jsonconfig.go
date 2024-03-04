/*
@File   : jsonconfig.go
@Author : pan
@Time   : 2023-06-09 16:57:51
*/
package main

import (
	"encoding/json"
	// "fmt"
	"io"
	"os"
)

type JsonConfig struct {
	Name     string         `json:"name,omitempty"`
	Hello    map[string]int `json:"hello,omitempty"`
	Words    []string       `json:"words,omitempty"`
	Users    Users          `json:"Users,omitempty"`
	Username string         `json:"username,omitempty"`
	Password string         `json:"password,omitempty"`
}

type Users struct {
	Student Student `json:"student,omitempty"`
	Teacher Teacher `json:"teacher,omitempty"`
}

type Student struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type Teacher struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func ParseJson(filepath string) (*JsonConfig, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var jsonconfig JsonConfig
	err = json.Unmarshal(data, &jsonconfig)
	if err != nil {
		return nil, err
	}
	return &jsonconfig, nil
}

func WriteJson(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	config := JsonConfig{
		Name:     "pan01",
		Hello:    map[string]int{"second": 2, "statuscode": 300},
		Words:    []string{"head", "queue", "set"},
		Username: "admin",
		Password: "admin123",
		Users:    Users{Student: Student{Name: "pan02", Age: 27}, Teacher: Teacher{Name: "pan03", Age: 25}},
	}
	// data, err := json.MarshalIndent(config, "", "\t")
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	f.Write(data)
	return nil
}

func main() {
	// data, err := ParseJson("../../config.json")
	// if err != nil {
	// 	fmt.Printf("parse json err:%v", err)
	// }
	// fmt.Println(data)
	WriteJson("../../writeconfig.json")
}
