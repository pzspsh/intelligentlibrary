package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for _, value := range rand.Perm(3) {
		fmt.Println(value)
	}

}
