/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:45:15
*/
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 如果输入部分字面量是字符串，你则不必使用正则。
	s := "abc,def,ghi"
	r, err := regexp.Compile(`[^,]+`) // everything that is not a comma
	if err != nil {
		fmt.Println(err)
	}
	res := r.FindAllString(s, -1)
	// Prints [abc def ghi]
	fmt.Printf("%v", res)

	/*
		// strings 包里面的 Split 函数就是用来做这个的，而且语法更可读。
		s = "abc,def,ghi"
		res = strings.Split(s, ",")
		// Prints [abc def ghi]
		fmt.Printf("%v", res)
	*/

	// 使用 MatchString 函数可以在一个字符串里查找另一个字面量的字符串。
	s = "OttoFritzHermanWaldoKarlSiegfried"
	r1, err := regexp.Compile(`Waldo`)
	if err != nil {
		fmt.Println(err)
	}
	res1 := r1.MatchString(s)
	// Prints true
	fmt.Printf("%v", res1)

	/*
		// 但是使用 strings.Index 函数可以在字串中获取匹配到子串的索引。当不匹配时则返回的索引为-1。
		s := "OttoFritzHermanWaldoKarlSiegfried"
		res:= strings.Index(s, "Waldo")
		// Prints true
		fmt.Printf("%v", res != -1)
	*/

	// 每当你读一些来自文件或是用户的文本时，你可能都想忽略那些句子开头和末尾的空格。
	s = "  Institute of Experimental Computer Science  "
	r2, err := regexp.Compile(`\s*(.*)\s*`)
	if err != nil {
		fmt.Println(err)
	}
	res2 := r2.FindStringSubmatch(s)
	// <Institute of Experimental Computer Science  >
	fmt.Printf("<%v>", res2[1])
	/*
		s := "  Institute of Experimental Computer Science  "
		// <Institute of Experimental Computer Science>
		fmt.Printf("<%v>", strings.TrimSpace(s))
	*/
}
