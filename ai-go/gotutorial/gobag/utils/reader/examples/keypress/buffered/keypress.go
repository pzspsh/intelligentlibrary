package main

import (
	"log"
	"sync"
	"time"

	"gotutorial/gobag/utils/reader"
	stringsutil "gotutorial/gobag/utils/strings"
)

func main() {
	stdr := reader.KeyPressReader{
		Timeout: time.Duration(5 * time.Second),
		Once:    &sync.Once{},
	}

	stdr.Start()
	defer stdr.Stop()

	for {
		data := make([]byte, stdr.BufferSize)
		n, err := stdr.Read(data)
		log.Println(n, err)

		if stringsutil.IsCTRLC(string(data)) {
			break
		}
	}
}
