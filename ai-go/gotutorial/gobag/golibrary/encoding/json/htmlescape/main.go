/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 16:31:34
*/
package main

import (
	"bytes"
	"encoding/json"
	"os"
)

func main() {
	var out bytes.Buffer
	json.HTMLEscape(&out, []byte(`{"Name":"<b>HTML content</b>"}`))
	out.WriteTo(os.Stdout)
}
