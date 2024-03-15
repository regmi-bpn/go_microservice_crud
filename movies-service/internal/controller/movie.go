package controller

import (
	"context"
	"github.com/google/uuid"
	"github.com/regmi-bpn/movie-common/pb"
	"github.com/regmi-bpn/movies-services/internal/entity"
	"github.com/regmi-bpn/movies-services/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Movie struct {
	repository repository.Movie
	pb.UnimplementedMovieServiceServer
}

func NewMovieController(repository repository.Movie) Movie {
	return Movie{
		repository: repository,
	}
}

func (c Movie) SaveMovie(ctx context.Context, req *pb.MovieSaveRequest) (*pb.MovieResponse, error) {
	mov, err := c.repository.Save(entity.Movie{
		ID:   uuid.New().String(),
		Name: req.GetName(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.MovieResponse{
		Id:   mov.ID,
		Name: mov.Name,
	}, nil
}
func (c Movie) UpdateMovie(ctx context.Context, req *pb.MovieUpdateRequest) (*pb.MovieResponse, error) {
	mov, err := c.repository.Update(req.GetId(), entity.Movie{
		Name: req.GetName(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.MovieResponse{
		Id:   mov.ID,
		Name: mov.Name,
	}, nil
}
func (c Movie) DeleteMovie(ctx context.Context, req *pb.MovieIdRequest) (*pb.EmptyMessage, error) {
	err := c.repository.Delete(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.EmptyMessage{}, nil
}
func (c Movie) GetMovie(ctx context.Context, req *pb.MovieIdRequest) (*pb.MovieResponse, error) {
	mov, err := c.repository.Get(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.MovieResponse{
		Id:   mov.ID,
		Name: mov.Name,
	}, nil
}
func (c Movie) GetMovies(ctx context.Context, req *pb.EmptyMessage) (*pb.MovieResponseList, error) {
	movs, err := c.repository.GetAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var list []*pb.MovieResponse

	for _, mov := range movs {
		list = append(list, &pb.MovieResponse{
			Id:   mov.ID,
			Name: mov.Name,
		})
	}

	return &pb.MovieResponseList{
		Movies: list,
	}, nil
}
