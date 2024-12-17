/*
@File   : main.go
@Author : pan
@Time   : 2024-12-17 10:55:28
*/
package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"strconv"

	. "gotutorial/examples/tcpsend/lib"

	"github.com/pkg/errors"
)

func client(ip string) error {
	cpData := ComplexData{
		N: 10,
		S: "测试string 数据",
		M: map[string]int{"A": 1, "B": 2},
		P: []byte("测试[]byte数据"),
		C: &ComplexData{
			N: 256,
			S: "Recursive structs? Piece of cake!",
			M: map[string]int{"01": 1, "10": 2, "11": 3},
		},
	}
	rw, err := Open(ip + Port)
	if err != nil {
		fmt.Println("客户端无法链接改地址：" + ip + Port)
		return err
	}
	n, err := rw.WriteString("string\n")
	if err != nil {
		return errors.Wrap(err, "Could not send the STRING request ("+strconv.Itoa(n)+" bytes written)")
	}
	n, err = rw.WriteString("Additional data.\n")
	if err != nil {
		return errors.Wrap(err, "Could not send additional STRING data ("+strconv.Itoa(n)+" bytes written)")
	}
	err = rw.Flush()
	if err != nil {
		return errors.Wrap(err, "Flush failed.")
	}

	// Read the reply.
	response, err := rw.ReadString('\n')
	if err != nil {
		return errors.Wrap(err, "Client: Failed to read the reply: '"+response+"'")
	}

	log.Println("STRING request: got a response:", response)

	log.Println("Send a struct as GOB:")
	log.Printf("Outer complexData struct: \n%#v\n", cpData)
	log.Printf("Inner complexData struct: \n%#v\n", cpData.C)
	enc := gob.NewEncoder(rw)
	n, err = rw.WriteString("gob\n")
	if err != nil {
		return errors.Wrap(err, "Could not write GOB data ("+strconv.Itoa(n)+" bytes written)")
	}
	err = enc.Encode(cpData)
	if err != nil {
		return errors.Wrapf(err, "Encode failed for struct: %#v", cpData)
	}
	err = rw.Flush()
	if err != nil {
		return errors.Wrap(err, "Flush failed.")
	}
	return nil
}

func main() {
	err := client("localhost")
	if err != nil {
		fmt.Println("Error:", errors.WithStack(err))
	}
}
