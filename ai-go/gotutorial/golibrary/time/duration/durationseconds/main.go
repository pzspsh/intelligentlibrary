package main

import (
	"fmt"
	"time"
)

func main() {
	m, _ := time.ParseDuration("1m30s")
	fmt.Printf("Take off in t-%.0f seconds.", m.Seconds())
}
