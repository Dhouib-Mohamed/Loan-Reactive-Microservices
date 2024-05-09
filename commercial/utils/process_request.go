package utils

import (
	"bytes"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"net/http"
)

func ProcessRequest(bytes_i []byte) {
	topic_succ := "request-loan-topic"
	topic_fail := "response-topic"
	producer, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:29092"})
	defer producer.Close()
	if bytes_i != nil {
		r, _ := http.Post("http://ocr-extraction:8080/ocr", "application/json", bytes.NewReader(bytes_i))
		if validate(r) {
			err := producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic_succ, Partition: kafka.PartitionAny},
				Value:          bytes_i,
			}, nil)
			fmt.Println("Commerial Verification Success")
			if err != nil {
				return
			}
		} else {
			err := producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic_fail, Partition: kafka.PartitionAny},
				Value:          bytes_i,
				Key:            []byte("fail"),
			}, nil)
			fmt.Println("Commercial Verification Failed")
			if err != nil {
				return
			}
		}
	}
}
func validate(r *http.Response) bool {
	// Validate Response
	return true
}
