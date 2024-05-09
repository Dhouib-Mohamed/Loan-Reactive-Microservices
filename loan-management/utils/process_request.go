package utils

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProcessRequest(bytes []byte) {
	topic := "user-request-topic"
	producer, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:29092"})
	defer producer.Close()
	SaveDB(bytes)
	if bytes != nil {
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          bytes,
		}, nil)
		if err != nil {
			return
		}
		Notify(bytes, "Request is Being Processed")
	}
	Notify(bytes, "Unexpected Error Occurred In the request")
}
