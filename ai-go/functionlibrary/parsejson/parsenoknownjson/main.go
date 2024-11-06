/*
@File   : main.go
@Author : pan
@Time   : 2024-11-06 10:55:27
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

func extractKeyValuePairs(data interface{}, parentKey string, result map[string]interface{}) {
	/*
		extractKeyValuePairs 递归地提取所有键值对
	*/
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)
	switch typ.Kind() {
	case reflect.Map: // 处理映射类型
		for _, k := range val.MapKeys() {
			v := val.MapIndex(k)
			mapKey := k.String()
			newKey := mapKey
			if parentKey != "" {
				newKey = parentKey + "." + mapKey
			}
			extractKeyValuePairs(v.Interface(), newKey, result) // 递归调用以处理映射的值（可能是另一个映射、切片或结构体）
		}
	case reflect.Slice, reflect.Array: // 处理切片或数组类型
		for i := 0; i < val.Len(); i++ {
			element := val.Index(i)
			newKey := fmt.Sprintf("%s[%d]", parentKey, i)
			extractKeyValuePairs(element.Interface(), newKey, result) // 递归调用以处理切片的元素（可能是另一个映射、切片或结构体）
		}
	case reflect.Struct: // 处理结构体类型
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldName := typ.Field(i).Name
			newKey := fieldName
			if parentKey != "" {
				newKey = parentKey + "." + fieldName
			}
			if field.CanInterface() { // 根据字段的类型递归调用或存储键值对
				result[newKey] = field.Interface()
			} else {
				extractKeyValuePairs(field.Interface(), newKey, result)
			}
		}
	default:
		if parentKey != "" { // 处理其他类型（如基本类型、指针等）
			result[parentKey] = data
		} else {
			/*
				如果 parentKey 为空，我们可能需要一个特殊的处理方式，
				或者简单地跳过非结构体、非映射、非切片的顶级元素。
				这里我们假设顶级元素应该直接存储（但这通常不是个好主意，因为它会丢失键信息）。
				更好的做法可能是抛出一个错误或记录一个警告。
			*/
			log.Printf("Unhandled type at top level: %T (value: %v)", data, data)
		}
	}
}
func main() {
	var err error
	var data map[string]interface{}           // 使用 map[string]interface{} 解析 JSON 数据
	var result = make(map[string]interface{}) // 存储所有键值对的结果
	file, err := os.Open("/path/data.json")   // 打开 JSON 文件
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file) // 读取文件内容
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	if err = json.Unmarshal(bytes, &data); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}

	extractKeyValuePairs(data, "", result) // 提取键值对
	for k, v := range result {             // 打印所有键值对
		fmt.Printf("Key: %s, Value: %v\n", k, v)
	}
}
