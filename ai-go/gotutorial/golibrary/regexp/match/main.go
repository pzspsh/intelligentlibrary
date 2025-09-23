package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
	fmt.Println(matched, err)
	matched, err = regexp.Match(`bar.*`, []byte(`seafood`))
	fmt.Println(matched, err)
	// matched, err = regexp.Match(`a(b`, []byte(`seafood`))
	matched, err = regexp.Match(`a(b)`, []byte(`seafood`))
	fmt.Println(matched, err)
}
