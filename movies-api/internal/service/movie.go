package service

import (
	"context"
	"github.com/regmi-bpn/movie-common/pb"
	commonTypes "github.com/regmi-bpn/movie-common/types"
	"github.com/regmi-bpn/movies-api/internal/publisher"
	"github.com/regmi-bpn/movies-api/types"
)

type MovieService interface {
	SaveMovie(request types.MovieRequest) (*types.MovieResponse, error)
	UpdateMovie(id string, request types.MovieRequest) (*types.MovieResponse, error)
	DeleteMovie(id string) error
	GetMovie(id string) (*types.MovieResponse, error)
	GetMovies() ([]types.MovieResponse, error)
	AddRating(id string, request types.RatingRequest) error
}

type Movie struct {
	client      pb.MovieServiceClient
	publisher   publisher.Publisher
	ratingTopic string
}

func NewMovieService(client pb.MovieServiceClient, publisher publisher.Publisher, ratingTopic string) MovieService {
	return Movie{
		client:      client,
		publisher:   publisher,
		ratingTopic: ratingTopic,
	}
}

func (c Movie) SaveMovie(request types.MovieRequest) (*types.MovieResponse, error) {
	r, err := c.client.SaveMovie(context.Background(), &pb.MovieSaveRequest{
		Name: request.Name,
	})
	if err != nil {
		return nil, err
	}
	return &types.MovieResponse{
		Id:   r.GetId(),
		Name: r.GetName(),
	}, nil
}

func (c Movie) UpdateMovie(id string, request types.MovieRequest) (*types.MovieResponse, error) {
	r, err := c.client.UpdateMovie(context.Background(), &pb.MovieUpdateRequest{
		Id:   id,
		Name: request.Name,
	})
	if err != nil {
		return nil, err
	}
	return &types.MovieResponse{
		Id:   r.GetId(),
		Name: r.GetName(),
	}, nil

}

func (c Movie) DeleteMovie(id string) error {
	_, err := c.client.DeleteMovie(context.Background(), &pb.MovieIdRequest{
		Id: id,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c Movie) GetMovie(id string) (*types.MovieResponse, error) {
	r, err := c.client.GetMovie(context.Background(), &pb.MovieIdRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return &types.MovieResponse{
		Id:   r.GetId(),
		Name: r.GetName(),
		Rating: types.MovieRatingResponse{
			Like:    r.Rating.Like,
			Dislike: r.Rating.Dislike,
		},
	}, nil
}

func (c Movie) GetMovies() ([]types.MovieResponse, error) {
	r, err := c.client.GetMovies(context.Background(), &pb.EmptyMessage{})
	if err != nil {
		return nil, err
	}
	var moviesResponse []types.MovieResponse

	for _, movie := range r.GetMovies() {
		moviesResponse = append(moviesResponse, types.MovieResponse{
			Id:   movie.GetId(),
			Name: movie.GetName(),
		})
	}

	return moviesResponse, nil
}

func (c Movie) AddRating(id string, request types.RatingRequest) error {
	return c.publisher.Publish(commonTypes.RatingConsumeMessage{
		MovieId: id,
		Like:    request.Like,
	}, c.ratingTopic)
}
