package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/regmi-bpn/movie-common/pb"
	"github.com/regmi-bpn/movies-services/internal/entity"
	"github.com/regmi-bpn/movies-services/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MovieService interface {
	SaveMovie(req *pb.MovieSaveRequest) (*pb.MovieResponse, error)
	UpdateMovie(req *pb.MovieUpdateRequest) (*pb.MovieResponse, error)
	DeleteMovie(req *pb.MovieIdRequest) (*pb.EmptyMessage, error)
	GetMovie(req *pb.MovieIdRequest) (*pb.MovieResponse, error)
	GetMovies(req *pb.EmptyMessage) (*pb.MovieResponseList, error)
}

type movieService struct {
	repository repository.MovieRepository
	client     pb.RatingServiceClient
}

func (m movieService) SaveMovie(req *pb.MovieSaveRequest) (*pb.MovieResponse, error) {
	mov, err := m.repository.Save(&entity.Movie{
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

func (m movieService) UpdateMovie(req *pb.MovieUpdateRequest) (*pb.MovieResponse, error) {
	mov, err := m.repository.Update(req.GetId(), &entity.Movie{
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

func (m movieService) DeleteMovie(req *pb.MovieIdRequest) (*pb.EmptyMessage, error) {
	err := m.repository.DeleteMovie(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.EmptyMessage{}, nil
}

func (m movieService) GetMovie(req *pb.MovieIdRequest) (*pb.MovieResponse, error) {
	mov, err := m.repository.Get(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	rating, err := m.client.GetRating(context.Background(), &pb.RatingRequest{MovieId: mov.ID})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.MovieResponse{
		Id:     mov.ID,
		Name:   mov.Name,
		Rating: &pb.MovieRating{Like: rating.Like, Dislike: rating.Dislike},
	}, nil
}

func (m movieService) GetMovies(req *pb.EmptyMessage) (*pb.MovieResponseList, error) {
	movs, err := m.repository.GetAll()
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

func NewMovieService(repository repository.MovieRepository, client pb.RatingServiceClient) MovieService {
	return &movieService{
		repository: repository,
		client:     client,
	}

}
