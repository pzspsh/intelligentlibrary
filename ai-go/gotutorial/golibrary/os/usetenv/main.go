/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 18:45:19
*/
package main

import (
	"os"
)

func main() {
	os.Setenv("TMPDIR", "/my/tmp")
	defer os.Unsetenv("TMPDIR")
}
