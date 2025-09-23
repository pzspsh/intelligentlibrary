package output

import "fmt"

var ResultChan = make(chan Result, 10000)

type Result struct {
}

func ResultBack() {
	for {
		if data, ok := <-ResultChan; ok {
			fmt.Println(data)
		}
	}
}
