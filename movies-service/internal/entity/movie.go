package entity

type Movie struct {
	ID   string `gorm:"type:char(36);primaryKey;unique;column:id" json:"id"`
	Name string `gorm:"type:varchar(100);not null;column:name" json:"name"`
}

func (Movie) TableName() string {
	return "movies"
}
