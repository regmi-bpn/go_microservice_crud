package main

import (
	"fmt"
	"github.com/regmi-bpn/movie-common/pb"
	"github.com/regmi-bpn/movies-services/internal/config"
	"github.com/regmi-bpn/movies-services/internal/controller"
	"github.com/regmi-bpn/movies-services/internal/repository"
	"github.com/regmi-bpn/movies-services/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	port := os.Getenv("MOVIE_GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	db := config.InitializeDatabase(config.Database{
		Host:     os.Getenv("MOVIE_DB_HOST"),
		Port:     os.Getenv("MOVIE_DB_PORT"),
		Username: os.Getenv("MOVIE_DB_USERNAME"),
		Password: os.Getenv("MOVIE_DB_PASSWORD"),
		Schema:   os.Getenv("MOVIE_DB_SCHEMA"),
	})

	s := grpc.NewServer()
	repo := repository.NewMovieRepository(db)
	serv := service.NewMovieService(repo)
	con := controller.NewMovieController(serv)
	pb.RegisterMovieServiceServer(s, con)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
