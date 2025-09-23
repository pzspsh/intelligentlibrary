package main

// import "C"
// import "runtime/cgo"

// //export MyGoPrint
// func MyGoPrint(handle C.uintptr_t) {
// 	h := cgo.Handle(handle)
// 	val := h.Value().(string)
// 	println(val)
// 	h.Delete()
// }

// func main() {
// 	val := "hello Go"
// 	C.myprint(C.uintptr_t(cgo.NewHandle(val)))
// 	// Output: hello Go
// }
