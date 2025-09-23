package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"function/traffichijacking/packet"
)

var (
	eventCh     = make(chan any, 1024)
	ctx, cancel = context.WithCancel(context.Background())
)

func main() {
	go shutdown()
	srv := packet.NewPacketHandle(ctx, "ens33", eventCh)
	srv.SetBpf("tcp port 80")     // option
	srv.SetEventHandle(5, handle) // option
	srv.SetPromisc(true)          // option
	srv.SetFlushTime(time.Minute) // option
	if err := srv.Listen(); err != nil {
		log.Println(err.Error())
	}
}

func handle(req *http.Request, resp *http.Response) {
	log.Printf("request uri: %s, response status: %v", req.RequestURI, resp.Status)
}

func shutdown() {
	time.Sleep(time.Second * 10)
	cancel()
}
