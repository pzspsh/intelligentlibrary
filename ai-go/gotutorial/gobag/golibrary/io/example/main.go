/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:54:52
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Copy1() {
	r := strings.NewReader("wang fan")
	buf := make([]byte, 20)
	r.Read(buf)
	fmt.Printf("string(buf): %v\n", string(buf))
}

func Copy2() {
	r := strings.NewReader("hello world") // 输出在控制台
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}

func Copy3() {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)
	_, err1 := io.CopyBuffer(os.Stdout, r1, buf)
	if err1 != nil {
		log.Fatal(err1)
	}
	_, err2 := io.CopyBuffer(os.Stdout, r2, buf)
	if err2 != nil {
		log.Fatal(err2)
	}
}

func Copy4() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 6)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

func Copy5() {
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	r := io.MultiReader(r1, r2, r3)
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func Copy6() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	var buf1, buf2 bytes.Buffer
	w := io.MultiWriter(&buf1, &buf2)
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("buf1.String(): %v\n", buf1.String())
	fmt.Printf("buf2.String(): %v\n", buf2.String())
}

func Copy7() {
	pr, pw := io.Pipe()
	go func() {
		fmt.Fprint(pw, "some io.Reader stream to be read\n")
		pw.Close()
	}()
	if _, err := io.Copy(os.Stdout, pr); err != nil {
		log.Fatal(err)
	}
}

func Copy8() {
	r := strings.NewReader("Go is general-purpose language designed with systems programming in mind.")
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("b: %v\n", string(b))
}

func Copy9() {
	// 拆一部分
	r := strings.NewReader("Go is general-purpose language designed with systems programming in mind.")
	s := io.NewSectionReader(r, 5, 17)
	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Copy1()
	// Copy2()
	// Copy3()
	// Copy4()
	// Copy5()
	// Copy6()
	// Copy7()
	// Copy8()
	Copy9()
}
