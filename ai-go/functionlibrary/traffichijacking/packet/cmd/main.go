package main

import (
	"context"
	"fmt"
	"log"

	"function/traffichijacking/packet"
)

var eventCh = make(chan any, 1024)

func main() {
	// device :=
	// addrs, _ := net.Interfaces()
	// fmt.Println("addrs: ", addrs[2].Name)
	device := "/Device/NPF_{B7E01C74-876F-44B9-95FF-13BAB4AB8D9E}"
	go handle()
	if err := packet.NewPacketHandle(context.Background(), device, eventCh).Listen(); err != nil {
		log.Println(err.Error())
	}
}

func handle() {
	for i := range eventCh {
		data := i.(packet.Event)
		fmt.Println(data.Req)
		fmt.Println("============================================")
		log.Printf("request uri: %s, response status: %v", data.Req.RequestURI, data.Resp.Status)
		fmt.Println("============================================")
		fmt.Println(data.Resp)
	}
}
