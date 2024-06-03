/*
@File   : consumer.go
@Author : pan
@Time   : 2024-06-03 10:40:16
*/
package server

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
)

func (p *PulsarOptions) Consumer(receive chan []byte) error {
	consumer, err := p.client.Subscribe(pulsar.ConsumerOptions{
		Topic:            p.Topic,
		SubscriptionName: p.Subscribe,
		Type:             pulsar.Shared,
	})
	if err != nil {
		return err
	}
	defer consumer.Close()
	ctx := context.Background()
	for {
		data, err := consumer.Receive(ctx)
		if err != nil {
			return err
		} else {
			receive <- data.Payload()
		}
		consumer.Ack(data)
	}
}
