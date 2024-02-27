/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:53:11
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

//os包中关于文件的操作函数
func main() {
	//创建文件，返回一个文件指针
	f3, _ := os.Create("./3.txt")
	defer f3.Close()

	//以读写方式打开文件，如果不存在则创建文件，等同于上面os.Create
	f4, _ := os.OpenFile("./4.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f4.Close()

	//打开文件，返回文件指针
	f1, _ := os.Open("./1.txt")
	defer f1.Close()

	//修改文件权限，类似os.chmod
	f1.Chmod(0777)

	//修改文件所有者，类似os.chown
	f1.Chown(0, 0)

	//返回文件的句柄，通过NewFile创建文件需要文件句柄
	fmt.Println(f1.Fd())

	//从文件中读取数据
	buf := make([]byte, 128)
	//read每次读取数据到buf中
	for n, _ := f1.Read(buf); n != 0; n, _ = f1.Read(buf) {
		fmt.Println(string(buf[:n]))
	}

	//向文件中写入数据
	for i := 0; i < 5; i++ {
		f3.Write([]byte("写入数据" + strconv.Itoa(i) + "\r\n"))
	}

	//返回一对关联的文件对象
	//从r中可以读取到从w写入的数据
	r, w, _ := os.Pipe()
	//向w中写入字符串
	w.WriteString("写入w")
	buf2 := make([]byte, 128)
	//从r中读取数据
	n, _ := r.Read(buf)
	fmt.Println(string(buf2[:n]))

	//改变工作目录
	os.Mkdir("a", os.ModePerm)
	dir, _ := os.Open("a")
	//改变工作目录到dir，dir必须为一个目录
	dir.Chdir()
	fmt.Println(os.Getwd())

	//读取目录的内容，返回一个FileInfo的slice
	//参数大于0，最多返回n个FileInfo
	//参数小于等于0，返回所有FileInfo
	fi, _ := dir.Readdir(-1)
	for _, v := range fi {
		fmt.Println(v.Name())
	}

	//读取目录中文件对象的名字
	names, _ := dir.Readdirnames(-1)
	fmt.Println(names)

	//获取文件的详细信息，返回FileInfo结构
	fi3, _ := f3.Stat()
	//文件名
	fmt.Println(fi3.Name())
	//文件大小
	fmt.Println(fi3.Size())
	//文件权限
	fmt.Println(fi3.Mode())
	//文件修改时间
	fmt.Println(fi3.ModTime())
	//是否是目录
	fmt.Println(fi3.IsDir())
}
