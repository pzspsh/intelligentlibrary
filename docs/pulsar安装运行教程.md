### pulsar安装运行教程
```
docker run -it -p 6650:6650 -p 8080:8080 --mount source=pulsardata,target=/pulsar/data --mount source=pulsarconf,target=/pulsar/conf apachepulsar/pulsar:2.11.1 bin/pulsar standalone
```

```go
consumer.go
import (
	"context"
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"
)

func ReceiveTask(task string) {
	//实例化Pulsar client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://ip:port", // xx.xx.xx.xx代表Pulsar IP
	})
	if err != nil {
		fmt.Println("ReceiveTask Error: ", err)
	}
	//使用client对象实例化consumer
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "my-topic",
		SubscriptionName: "sub-demo",
		Type:             pulsar.Shared,
	})
	if err != nil {
		fmt.Println("consumer Error: ", err)
	}
	defer consumer.Close()
	ctx := context.Background()
	//无限循环监听topic
	for {
		taskData, err := consumer.Receive(ctx)
		if err != nil {
			fmt.Println("taskData Error: ", err)
		} else {
			if taskData != nil {
				task = string(taskData.Payload())
				fmt.Println(task)
			}
		}
		consumer.Ack(taskData)
	}
}

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func PushResult() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar+ssl://10.0.35.74:6651",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		fmt.Printf("Could not instantiate Pulsar client: %v", err)
	}

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})

	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		aa := map[string]string{
			"target":   "http://xx.xx.xx.xx:port",
			"Count":    strconv.Itoa(i),
		}
		b, _ := json.Marshal(aa)
		msg := pulsar.ProducerMessage{
			Payload: []byte(string(b)),
		}

		if err, _ := producer.Send(context.Background(), &msg); err != nil {
			fmt.Printf("Producer could not send message:%v", err)
		}

		defer producer.Close()

		if err != nil {
			fmt.Println("Failed to publish message", err)
		}
		fmt.Println("Published message")
	}
}

```
