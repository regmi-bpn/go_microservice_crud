package config

import (
	"fmt"
	"github.com/regmi-bpn/rating-service/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	Username string
	Password string
	Host     string
	Port     string
	Schema   string
}

func InitializeDatabase(config Database) *gorm.DB {
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Username, config.Password, config.Host, config.Port, config.Schema)
	db, err := gorm.Open(mysql.Open(datasource))
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
	}

	err = db.AutoMigrate(&entity.Rating{})
	if err != nil {
		log.Fatalf("error migrating database: %v", err)
	}

	return db
}
