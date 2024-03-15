package main

import (
	"fmt"
	"github.com/regmi-bpn/movie-common/pb"
	"github.com/regmi-bpn/movies-api/internal/controller"
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

	c := pb.NewMovieServiceClient(conn)
	movieService := service.NewMovieService(c)
	movieController := controller.NewMovieController(movieService)
	r := routes.InitRoutes(movieController)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Error listening on port %s: %v", port, err)
	}
}
