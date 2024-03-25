package service

import (
	"errors"
	"github.com/regmi-bpn/movie-common/pb"
	"github.com/regmi-bpn/rating-service/internal/entity"
	"github.com/regmi-bpn/rating-service/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type RatingService interface {
	GetRating(req *pb.RatingRequest) (*pb.RatingResponse, error)
	SaveRating(movieId string, like bool) error
}

type ratingService struct {
	repository repository.RatingRepository
}

func (r ratingService) SaveRating(movieId string, like bool) error {
	rat, err := r.repository.Get(movieId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if rat == nil {
		return r.createRating(movieId, like)
	}
	if like {
		rat.Like++
	} else {
		rat.Dislike++
	}
	_, err = r.repository.Update(movieId, *rat)
	if err != nil {
		return err
	}
	return nil
}

func (r ratingService) createRating(movieId string, like bool) error {
	rat := &entity.Rating{
		MovieID: movieId,
		Like:    0,
		Dislike: 0,
	}
	if like {
		rat.Like++
	} else {
		rat.Dislike++
	}

	_, err := r.repository.Save(*rat)
	if err != nil {
		return err
	}
	return nil

}

func (r ratingService) GetRating(req *pb.RatingRequest) (*pb.RatingResponse, error) {
	rat, err := r.repository.Get(req.GetMovieId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.RatingResponse{MovieId: rat.MovieID, Like: rat.Like, Dislike: rat.Dislike}, nil

}

func NewRatingService(repository repository.RatingRepository) RatingService {
	return ratingService{repository: repository}

}
