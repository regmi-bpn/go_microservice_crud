package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/regmi-bpn/movie-common/pb"
	"github.com/regmi-bpn/movies-api/internal/controller"
	"github.com/regmi-bpn/movies-api/internal/publisher"
	"github.com/regmi-bpn/movies-api/internal/routes"
	"github.com/regmi-bpn/movies-api/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {
	port := os.Getenv("MOVIE_REST_PORT")
	conn, err := grpc.Dial(os.Getenv("MOVIE_GRPC_HOST"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to grpc server: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	config := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("RATING_KAFKA_SERVER"),
	}
	topic := os.Getenv("RATING_KAFKA_TOPIC")

	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	pub := publisher.NewPublisher(producer)

	c := pb.NewMovieServiceClient(conn)
	movieService := service.NewMovieService(c, pub, topic)
	movieController := controller.NewMovieController(movieService)
	r := routes.InitRoutes(movieController)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Error listening on port %s: %v", port, err)
	}
}
