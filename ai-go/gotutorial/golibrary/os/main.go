/*
@File   : main.go
@Author : pan
@Time   : 2023-12-07 16:50:45
*/
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// os包中的一些常用函数
func main() {
	//获取主机名
	fmt.Println(os.Hostname())

	//获取当前目录
	fmt.Println(os.Getwd())

	//获取用户ID
	fmt.Println(os.Getuid())

	//获取有效用户ID
	fmt.Println(os.Geteuid())

	//获取组ID
	fmt.Println(os.Getgid())

	//获取有效组ID
	fmt.Println(os.Getegid())

	//获取进程ID
	fmt.Println(os.Getpid())

	//获取父进程ID
	fmt.Println(os.Getppid())

	//获取环境变量的值
	fmt.Println(os.Getenv("GOPATH"))

	//设置环境变量的值
	os.Setenv("TEST", "test")

	//改变当前工作目录
	os.Chdir("C:/")
	fmt.Println(os.Getwd())

	//创建文件
	f1, _ := os.Create("./1.txt")
	defer f1.Close()

	//修改文件权限
	if err := os.Chmod("./1.txt", 0777); err != nil {
		fmt.Println(err)
	}

	//修改文件所有者
	if err := os.Chown("./1.txt", 0, 0); err != nil {
		fmt.Println(err)
	}

	//修改文件的访问时间和修改时间
	os.Chtimes("./1.txt", time.Now().Add(time.Hour), time.Now().Add(time.Hour))

	//获取所有环境变量
	fmt.Println(strings.Join(os.Environ(), "\r\n"))

	//把字符串中带${var}或$var替换成指定指符串
	fmt.Println(os.Expand("${1} ${2} ${3}", func(k string) string {
		mapp := map[string]string{
			"1": "111",
			"2": "222",
			"3": "333",
		}
		return mapp[k]
	}))

	//创建目录
	os.Mkdir("abc", os.ModePerm)

	//创建多级目录
	os.MkdirAll("abc/d/e/f", os.ModePerm)

	//删除文件或目录
	os.Remove("abc/d/e/f")

	//删除指定目录下所有文件
	os.RemoveAll("abc")

	//重命名文件
	os.Rename("./2.txt", "./2_new.txt")

	//判断是否为同一文件
	//unix下通过底层结构的设备和索引节点是否相同来判断
	//其他系统可能是通过文件绝对路径来判断
	fs1, _ := f1.Stat()
	f2, _ := os.Open("./1.txt")
	fs2, _ := f2.Stat()
	fmt.Println(os.SameFile(fs1, fs2))

	//返回临时目录
	fmt.Println(os.TempDir())
}
