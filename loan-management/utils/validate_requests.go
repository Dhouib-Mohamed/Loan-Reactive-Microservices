package utils

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

func ValidateRequests() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "foo",
		"auto.offset.reset": "smallest"})
	if err != nil {
		panic(err)
	}
	err = consumer.SubscribeTopics([]string{"response-topic"}, nil)
	if err != nil {
		panic(err)
	}
	run := true
	for run {
		ev, err := consumer.ReadMessage(100 * time.Millisecond)
		if err != nil {
			continue
		}
		if isRequestValidated(ev.Value) {
			Notify(ev.Value, "Request is Validated")
		} else {
			Notify(ev.Value, "Request is Invalid")
		}
	}
}

func isRequestValidated(bytes []byte) bool {
	// Validate Request
	return true
}
