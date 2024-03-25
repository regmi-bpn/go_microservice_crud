package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/regmi-bpn/movie-common/pb"
	"github.com/regmi-bpn/rating-service/internal/config"
	"github.com/regmi-bpn/rating-service/internal/consumer"
	"github.com/regmi-bpn/rating-service/internal/controller"
	"github.com/regmi-bpn/rating-service/internal/repository"
	"github.com/regmi-bpn/rating-service/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	port := os.Getenv("RATING_GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	db := config.InitializeDatabase(config.Database{
		Host:     os.Getenv("RATING_DB_HOST"),
		Port:     os.Getenv("RATING_DB_PORT"),
		Username: os.Getenv("RATING_DB_USERNAME"),
		Password: os.Getenv("RATING_DB_PASSWORD"),
		Schema:   os.Getenv("RATING_DB_SCHEMA"),
	})

	kafkaConf := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("RATING_KAFKA_SERVER"),
		"group.id":          os.Getenv("RATING_KAFKA_GROUP"),
		"auto.offset.reset": "earliest",
	}

	kafkaConsumer, err := kafka.NewConsumer(kafkaConf)
	if err != nil {
		panic(err)
	}
	err = kafkaConsumer.SubscribeTopics([]string{os.Getenv("RATING_KAFKA_TOPIC")}, nil)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	repo := repository.NewRatingRepository(db)
	serv := service.NewRatingService(repo)
	con := controller.NewRatingController(serv)
	pb.RegisterRatingServiceServer(s, con)

	cons := consumer.NewRatingConsumer(kafkaConsumer, serv)
	go func(ratingConsumer *consumer.RatingConsumer) {
		ratingConsumer.Consume()
	}(cons)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
