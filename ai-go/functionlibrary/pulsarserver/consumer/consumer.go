/*
@File   : consumer.go
@Author : pan
@Time   : 2023-06-12 15:31:49
*/
package main

import (
	"context"
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"
)

func ConsumerCert(receive chan []byte) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		// URL: "pulsar://127.0.0.1:6650",
		URL:                        "pulsar+ssl://127.0.0.1:6651",
		TLSTrustCertsFilePath:      `证书路径/ca.cert.pem`,
		TLSAllowInsecureConnection: true,
	})
	if err != nil {
		fmt.Println(err)
	}
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "my-topic",
		SubscriptionName: "sub-demo",
		Type:             pulsar.Shared,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer consumer.Close()
	ctx := context.Background()
	for {
		taskData, err := consumer.Receive(ctx)
		if err != nil {
			fmt.Println(err)
		} else {
			if taskData != nil {
				receive <- taskData.Payload()
				fmt.Println(string(taskData.Payload()))
			}
		}
		consumer.Ack(taskData)
	}
}

func Consumer(receive chan any) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar0://127.0.0.1:6650",
	})
	if err != nil {
		fmt.Println(err)
	}
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "my-topic",
		SubscriptionName: "sub-demo",
		Type:             pulsar.Shared,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer consumer.Close()
	ctx := context.Background()
	for {
		taskData, err := consumer.Receive(ctx)
		if err != nil {
			fmt.Println(err)
		} else {
			if taskData != nil {
				receive <- taskData.Payload()
				fmt.Println(string(taskData.Payload()))
			}
		}
		consumer.Ack(taskData)
	}
}

func main() {

}
