/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 12:00:30
*/
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	r := bytes.NewReader([]byte("ABCDEFGHIJKLMN IIIIII LLLLLLLL SSSSSS"))

	fmt.Println(r.Len())  // 37
	fmt.Println(r.Size()) // 37

	tmp := make([]byte, 5)
	n, _ := r.Read(tmp)
	fmt.Println(string(tmp[:n]))   // ABCDE
	fmt.Println(r.Len(), r.Size()) // 32 37

	fmt.Println(r.ReadByte())      // 70 <nil> // F
	fmt.Println(r.ReadRune())      // 71 1 <nil>
	fmt.Println(r.Len(), r.Size()) // 30 37

	b := []byte("III") // cap 3
	n, _ = r.ReadAt(b, 1)
	fmt.Println(string(b), n) // BCD 3

	r.Reset([]byte("Hi,My god"))
	fmt.Println(r.Len(), r.Size()) // 9 9

	r.WriteTo(os.Stdout) // Hi,My god
}
