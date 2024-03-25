package controller

import (
	"context"
	"github.com/regmi-bpn/movie-common/pb"
	"github.com/regmi-bpn/movies-services/internal/service"
)

type MovieController struct {
	pb.UnimplementedMovieServiceServer
	service service.MovieService
}

func NewMovieController(service service.MovieService) *MovieController {
	return &MovieController{
		service: service,
	}
}

func (c MovieController) SaveMovie(ctx context.Context, req *pb.MovieSaveRequest) (*pb.MovieResponse, error) {
	mov, err := c.service.SaveMovie(req)
	if err != nil {
		return nil, err
	}

	response := &pb.MovieResponse{
		Id:   mov.Id,
		Name: mov.Name,
	}

	return response, nil
}
func (c MovieController) UpdateMovie(ctx context.Context, req *pb.MovieUpdateRequest) (*pb.MovieResponse, error) {
	mov, err := c.service.UpdateMovie(req)

	if err != nil {
		return nil, err

	}

	response := &pb.MovieResponse{Id: mov.Id, Name: mov.Name}
	return response, nil
}
func (c MovieController) DeleteMovie(ctx context.Context, req *pb.MovieIdRequest) (*pb.EmptyMessage, error) {
	_, err := c.service.DeleteMovie(req)

	if err != nil {
		return nil, err
	}
	return nil, err

}
func (c MovieController) GetMovie(ctx context.Context, req *pb.MovieIdRequest) (*pb.MovieResponse, error) {
	mov, err := c.service.GetMovie(req)

	if err != nil {
		return nil, err

	}

	return mov, nil

}
func (c MovieController) GetMovies(ctx context.Context, req *pb.EmptyMessage) (*pb.MovieResponseList, error) {

	moves, err := c.service.GetMovies(req)

	if err != nil {
		return nil, err

	}

	return moves, nil

}
