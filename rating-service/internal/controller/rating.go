package controller

import (
	"context"
	"github.com/regmi-bpn/movie-common/pb"
	"github.com/regmi-bpn/rating-service/internal/service"
)

type RatingController struct {
	pb.UnimplementedRatingServiceServer
	service service.RatingService
}

func NewRatingController(service service.RatingService) *RatingController {
	return &RatingController{
		service: service,
	}

}

func (c RatingController) GetRating(ctx context.Context, req *pb.RatingRequest) (*pb.RatingResponse, error) {
	mov, err := c.service.GetRating(req)

	if err != nil {
		return nil, err

	}

	return mov, nil
}
