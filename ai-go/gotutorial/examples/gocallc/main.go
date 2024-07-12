/*
@File   : main.go
@Author : pan
@Time   : 2024-07-12 14:55:33
*/
// Go调用C
package main

// #include <stdio.h>
// void printHelloFromC();
import "C"
import "fmt"

func main() {
	C.printHelloFromC()
	fmt.Println("Hello from Go!")
}
