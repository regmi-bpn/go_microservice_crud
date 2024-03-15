package repository

import (
	"github.com/regmi-bpn/movies-services/internal/entity"
	"gorm.io/gorm"
)

type Movie struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) Movie {
	return Movie{
		db: db,
	}
}

func (u Movie) Get(id string) (*entity.Movie, error) {
	var movie entity.Movie
	err := u.db.Model(&entity.Movie{}).Where("id = ?", id).First(&movie).Error
	return &movie, err
}

func (u Movie) GetAll() ([]entity.Movie, error) {
	var movies []entity.Movie
	err := u.db.Model(&entity.Movie{}).Find(&movies).Error
	return movies, err
}

func (u Movie) Save(movie entity.Movie) (*entity.Movie, error) {
	err := u.db.Model(&entity.Movie{}).Create(&movie).Error
	return &movie, err
}

func (u Movie) Update(id string, movie entity.Movie) (*entity.Movie, error) {
	err := u.db.Model(&entity.Movie{}).Where("id = ?", id).Updates(&movie).Error
	movie.ID = id
	return &movie, err
}

func (u Movie) Delete(id string) error {
	return u.db.Model(&entity.Movie{}).Where("id = ?", id).Delete(&entity.Movie{}).Error
}
