package publisher

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Publisher struct {
	producer *kafka.Producer
}

func NewPublisher(producer *kafka.Producer) Publisher {
	return Publisher{
		producer: producer,
	}
}

func (p Publisher) Publish(payload interface{}, topic string) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)
	if err != nil {
		return err
	}
	return nil
}
