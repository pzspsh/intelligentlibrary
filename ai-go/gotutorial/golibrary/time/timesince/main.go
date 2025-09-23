package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now().Add(-2 * time.Hour)
	timeDiff := time.Since(startTime)
	fmt.Printf("time diff: %s\n", timeDiff)
}
