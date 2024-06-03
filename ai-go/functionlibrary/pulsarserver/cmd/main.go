/*
@File   : main.go
@Author : pan
@Time   : 2024-06-03 11:04:26
*/
package main

import (
	"fmt"
	"function/pulsarserver/server"
)

func ConsumerRun() {
	var err error
	pulsaroption := &server.PulsarOptions{
		PulsarUrl: "pulsar+ssl://127.0.0.1:6651",
		CertPath:  "./ca.cert.pem", // 你生成的证书路径
		Topic:     "persistent://public/default/test",
		Subscribe: "test",
	}
	if pulsaroption, err = pulsaroption.PulsarConn(); err != nil {
		panic(err)
	}
	datas := make(chan []byte, 100)
	if err := pulsaroption.Consumer(datas); err != nil {
		panic(err)
	}
	for {
		if data, ok := <-datas; ok {
			fmt.Println(data)
		}
	}
}

func ProducerRun() {
	var err error
	pulsaroption := &server.PulsarOptions{
		PulsarUrl: "pulsar+ssl://127.0.0.1:6651",
		CertPath:  "./ca.cert.pem", // 你生成的证书路径
		Topic:     "persistent://public/default/test",
	}
	data := map[string]string{"aaa": "bbb"}
	if err = pulsaroption.Producer(data); err != nil {
		panic(err)
	}
}

func main() {
	ConsumerRun()
	// ProducerRun()
}
