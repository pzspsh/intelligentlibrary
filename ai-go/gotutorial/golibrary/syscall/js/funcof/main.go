/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 23:40:09
*/
package main

// import (
// 	"fmt"
// 	"syscall/js"
// )

// func main() {
// 	var cb js.Func
// 	cb = js.FuncOf(func(this js.Value, args []js.Value) any {
// 		fmt.Println("button clicked")
// 		cb.Release() // release the function if the button will not be clicked again
// 		return nil
// 	})
// 	js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)
// }
