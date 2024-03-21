package repository

import (
	"github.com/regmi-bpn/movies-services/internal/entity"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Get(id string) (*entity.Movie, error)
	GetAll() ([]entity.Movie, error)
	Save(movie *entity.Movie) (*entity.Movie, error)
	Update(id string, movie *entity.Movie) (*entity.Movie, error)
	DeleteMovie(id string) error
}

type movieRepository struct {
	db *gorm.DB
}

func (r *movieRepository) Get(id string) (*entity.Movie, error) {
	var movie entity.Movie
	err := r.db.Model(&entity.Movie{}).Where("id = ?", id).First(&movie).Error
	return &movie, err
}

func (r *movieRepository) GetAll() ([]entity.Movie, error) {
	var movies []entity.Movie
	err := r.db.Model(&entity.Movie{}).Find(&movies).Error
	return movies, err
}

func (r *movieRepository) Save(movie *entity.Movie) (*entity.Movie, error) {
	err := r.db.Model(&entity.Movie{}).Create(&movie).Error
	return movie, err
}

func (r *movieRepository) Update(id string, movie *entity.Movie) (*entity.Movie, error) {
	err := r.db.Model(&entity.Movie{}).Where("id = ?", id).Updates(&movie).Error
	movie.ID = id
	return movie, err
}

func (r *movieRepository) DeleteMovie(id string) error {
	return r.db.Model(&entity.Movie{}).Where("id = ?", id).Delete(&entity.Movie{}).Error

}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{
		db: db,
	}
}
