package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/regmi-bpn/movie-common/types"
	"github.com/regmi-bpn/rating-service/internal/service"
)

type RatingConsumer struct {
	consumer *kafka.Consumer
	service  service.RatingService
}

func NewRatingConsumer(consumer *kafka.Consumer, service service.RatingService) *RatingConsumer {
	return &RatingConsumer{
		consumer: consumer,
		service:  service,
	}
}

func (c RatingConsumer) Consume() {
	for {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Error while consuming message: %v\n", err)
			continue
		}
		var message types.RatingConsumeMessage
		err = json.Unmarshal(msg.Value, &message)
		if err != nil {
			fmt.Printf("Error while unmarshalling message: %v\n", err)
			continue
		}
		err = c.service.SaveRating(message.MovieId, message.Like)
		if err != nil {
			fmt.Printf("Error while saving rating: %v\n", err)
			continue
		}
		fmt.Printf("Rating saved successfully for movie %s\n", message.MovieId)
	}
}
