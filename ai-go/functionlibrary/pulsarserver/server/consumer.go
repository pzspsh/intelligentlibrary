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

type PulsarConfig struct {
	ReceivePulsar Pulsars `json:"ReceivePulsar,omitempty"`
	SendPulsar    Pulsars `json:"SendPulsar,omitempty"`
}

type Pulsars struct {
	LocalPulsar  Pulsar `json:"localpulsar,omitempty"`
	RemotePulsar Pulsar `json:"remotepulsar,omitempty"`
}

type Pulsar struct {
	Url     string   `json:"url,omitempty"`
	Topic   []string `json:"topic,omitempty"`
	SubName string   `json:"subname,omitempty"`
}

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
