package repository

import (
	"github.com/regmi-bpn/rating-service/internal/entity"
	"gorm.io/gorm"
)

type RatingRepository interface {
	Get(movieId string) (*entity.Rating, error)
	Save(rating entity.Rating) (*entity.Rating, error)
	Update(movieId string, rating entity.Rating) (*entity.Rating, error)
}

type ratingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) RatingRepository {
	return &ratingRepository{db: db}

}

func (r *ratingRepository) Get(movieId string) (*entity.Rating, error) {
	var rating entity.Rating
	err := r.db.Model(&entity.Rating{}).Where("movieId = ?", movieId).First(&rating).Error
	return &rating, err
}

func (r *ratingRepository) Save(rating entity.Rating) (*entity.Rating, error) {
	err := r.db.Model(&entity.Rating{}).Save(&rating).Error
	return &rating, err
}

func (r *ratingRepository) Update(movieId string, rating entity.Rating) (*entity.Rating, error) {
	err := r.db.Model(&entity.Rating{}).Where("movieId = ?", movieId).Updates(&rating).Error
	return &rating, err
}
