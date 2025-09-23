package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Add("cat sounds", "meow")
	v.Add("cat sounds", "mew")
	v.Add("cat sounds", "mau")
	fmt.Printf("%q\n", v.Get("cat sounds"))
	fmt.Printf("%q\n", v.Get("dog sounds"))

}
