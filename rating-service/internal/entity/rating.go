package entity

type Rating struct {
	MovieID string `gorm:"type:char(36);primaryKey;unique;column:movieId" json:"id"`
	Like    int64  `gorm:"type:int;not null;column:like" json:"like"`
	Dislike int64  `gorm:"type:int;not null;column:dislike" json:"dislike"`
}

func (Rating) TableName() string {
	return "rating"
}
