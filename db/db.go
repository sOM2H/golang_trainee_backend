package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sOM2H/golang_trainee_backend/model"
)

func New(dns string) *gorm.DB {
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		log.Println(err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
	)
}
