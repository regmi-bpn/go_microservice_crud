package service

import (
	"context"
	"github.com/regmi-bpn/movie-common/pb"
	"github.com/regmi-bpn/movies-api/types"
)

type Movie struct {
	client pb.MovieServiceClient
}

func NewMovieService(client pb.MovieServiceClient) Movie {
	return Movie{
		client: client,
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
