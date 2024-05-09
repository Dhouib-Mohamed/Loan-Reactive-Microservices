package utils

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

func RequestSubscribe() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "foo",
		"auto.offset.reset": "smallest"})
	if err != nil {
		panic(err)
	}
	err = consumer.SubscribeTopics([]string{"user-request-topic"}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Subscribed to user-request-topic")
	run := true
	for run {
		ev, err := consumer.ReadMessage(100 * time.Millisecond)
		if err != nil {
			continue
		}
		fmt.Println("Received message", string(ev.Value))
		go ProcessRequest(ev.Value)
	}
}
