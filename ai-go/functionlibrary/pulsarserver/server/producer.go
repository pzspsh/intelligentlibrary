package server

import (
	"context"
	"encoding/json"

	"github.com/apache/pulsar-client-go/pulsar"
)

func (p *PulsarOptions) Producer(data any) error {
	producer, err := p.client.CreateProducer(pulsar.ProducerOptions{
		Topic: p.Topic,
	})
	if err != nil {
		return err
	}
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	msg := pulsar.ProducerMessage{
		Payload: b,
	}
	if _, err = producer.Send(context.Background(), &msg); err != nil {
		return err
	}
	return nil
}
