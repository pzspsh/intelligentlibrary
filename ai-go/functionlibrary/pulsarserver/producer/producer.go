/*
@File   : producer.go
@Author : pan
@Time   : 2023-06-12 15:31:58
*/
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func ProducerCert() {
	Client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:                        "pulsar+ssl://127.0.0.1:6651",
		TLSTrustCertsFilePath:      "证书路径/ca.cert.pem",
		TLSAllowInsecureConnection: true,
		OperationTimeout:           30 * time.Second,
		ConnectionTimeout:          30 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	producer, err := Client.CreateProducer(pulsar.ProducerOptions{
		Topic: "",
	})
	if err != nil {
		fmt.Println(err)
	}
	var target = make(map[string]interface{})
	for i := 0; i < 9; i++ {
		target["target"] = "http://127.0.0.1:7001"
		b, _ := json.Marshal(target)
		msg := pulsar.ProducerMessage{
			Payload: []byte(string(b)),
		}
		if err, _ := producer.Send(context.Background(), &msg); err != nil {
			fmt.Println(err)
		}
	}
}

func Producer() {
	Client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://127.0.0.1:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	producer, err := Client.CreateProducer(pulsar.ProducerOptions{
		Topic: "",
	})
	if err != nil {
		fmt.Println(err)
	}
	var target = make(map[string]interface{})
	for i := 0; i < 9; i++ {
		target["target"] = "http://127.0.0.1:7001"
		b, _ := json.Marshal(target)
		msg := pulsar.ProducerMessage{
			Payload: []byte(string(b)),
		}
		if err, _ := producer.Send(context.Background(), &msg); err != nil {
			fmt.Println(err)
		}
	}
}

func Producer2(data any) { // 参数类型any,发生任何类型的数据
	Client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://127.0.0.1:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	producer, err := Client.CreateProducer(pulsar.ProducerOptions{
		Topic: "",
	})
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	msg := pulsar.ProducerMessage{
		Payload: b,
	}
	if err, _ := producer.Send(context.Background(), &msg); err != nil {
		fmt.Println(err)
	}
}

func main() {

}
