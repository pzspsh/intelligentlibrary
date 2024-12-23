/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 15:19:18
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"mime"
)

func main() {
	dec := new(mime.WordDecoder)
	header, err := dec.Decode("=?utf-8?q?=C2=A1Hola,_se=C3=B1or!?=")
	if err != nil {
		panic(err)
	}
	fmt.Println(header)

	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "x-case":
			// Fake character set for example.
			// Real use would integrate with packages such
			// as code.google.com/p/go-charset
			content, err := io.ReadAll(input)
			if err != nil {
				return nil, err
			}
			return bytes.NewReader(bytes.ToUpper(content)), nil
		default:
			return nil, fmt.Errorf("unhandled charset %q", charset)
		}
	}
	header, err = dec.Decode("=?x-case?q?hello!?=")
	if err != nil {
		panic(err)
	}
	fmt.Println(header)
}
