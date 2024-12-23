/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 11:50:30
*/
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"
)

// Go 1.21中的泛型基本语法
/*
要定义泛型函数或类型，可以使用类型 T关键字，后跟用方括号[]括起来的泛型形参的名称。
例如，要创建一个接受任意类型的slice并返回其第一个元素的泛型函数，可以这样定义:
*/
func First[T any](items []T) T {
	return items[0]
}

/*
在上面的例子中，[T any]表示类型参数T，它表示任意类型。any关键字表示T类型可以是任何有效类型。
然后，可以使用任何切片类型调用First函数，该函数将返回该切片的第一个元素
*/
func RunMain() {
	intSlice := []int{1, 2, 3, 4, 5}
	firstInt := First[int](intSlice) // returns 1

	println(firstInt)

	stringSlice := []string{"apple", "banana", "cherry"}
	firstString := First[string](stringSlice) // returns "apple"

	println(firstString)
}

// 在Go 1.21中具有各种类型的泛型
/*
在另一个示例中，让我们编写函数SumGenerics，它对各种数字类型(如int、int16、int32、int64、int8、float32和float64)执行加法操作。
*/
func SumGenerics[T int | int16 | int32 | int64 | int8 | float32 | float64](a, b T) T {
	return a + b
}

func RunMain2() {
	sumInt := SumGenerics[int](2, 3)
	sumFloat := SumGenerics[float32](2.5, 3.5)
	sumInt64 := SumGenerics[int64](10, 20)

	fmt.Println(sumInt)   // returns 5
	fmt.Println(sumFloat) // returns 6.0
	fmt.Println(sumInt64) // returns 30
}

// Go 1.21中具有任意数据类型的泛型
/*
泛型可以用于任意数据类型的序列化和反序列化，使用提供的序列化和反序列化函数:
*/
type Person struct {
	Name    string
	Age     int
	Address string
}

func Serialize[T any](data T) ([]byte, error) {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Deserialize[T any](b []byte) (T, error) {
	buffer := bytes.Buffer{}
	buffer.Write(b)
	decoder := gob.NewDecoder(&buffer)
	var data T
	err := decoder.Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

/*
在本例中，我们有两个通用函数Serialize和Deserialize，它们利用Go的gob包将任意数据类型转换为字节，反之亦然。
*/

func RunMain3() {
	person := Person{
		Name:    "John",
		Age:     30,
		Address: "123 Main St.",
	}

	serialized, err := Serialize(person)
	if err != nil {
		panic(err)
	}

	deserialized, err := Deserialize[Person](serialized)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %s, Age: %d, Address: %s", deserialized.Name, deserialized.Age, deserialized.Address)
	/*
			Output: Name: John, Age: 30, Address: 123 Main St.
			在上面的代码中，我们用一些数据创建了一个Person实例。然后使用Serialize函数将person对象转换为字节数组。稍后，
		使用Deserialize函数，将字节数组转换回Person对象。
			通过将序列化和反序列化函数定义为具有T any类型参数的泛型函数，我们可以序列化和反序列化任何支持使用gob包
		进行编码和解码的数据类型。
	*/
}

// 在Go中使用泛型和Validate函数自定义验证器
// 让我们用自定义验证器编写一个通用的Validate函数。
type Validator[T any] func(T) error

func Validate[T any](data T, validators ...Validator[T]) error {
	for _, validator := range validators {
		err := validator(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func StringNotEmpty(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return fmt.Errorf("string cannot be empty")
	}
	return nil
}

func IntInRange(num int, min, max int) error {
	if num < min || num > max {
		return fmt.Errorf("number must be between %d and %d", min, max)
	}
	return nil
}

func RunMain4() {
	person := Person{
		Name:    "John",
		Age:     30,
		Address: "123 Main St.",
	}

	err := Validate(person, func(p Person) error {
		return StringNotEmpty(p.Name)
	}, func(p Person) error {
		return IntInRange(p.Age, 0, 120)
	})

	if err != nil {
		println(err.Error())
		panic(err)
	}
	println("Person is valid")
}

type LoginForm struct {
	Username string
	Password string
}

func (f *LoginForm) Validate() error {
	return Validate(f,
		func(l *LoginForm) error {
			return StringNotEmpty(l.Username)
		},
		func(l *LoginForm) error {
			return StringNotEmpty(l.Password)
		},
	)
}

func RunMain5() {
	loginForm := LoginForm{
		Username: "John",
		Password: "123",
	}

	err := loginForm.Validate()
	if err != nil {
		println(err.Error())
		panic(err)
	}

	println("Login form is valid")
}

func main() {
	// RunMain()
	// RunMain2()
	// RunMain3()
	// RunMain4()
	RunMain5()
}
