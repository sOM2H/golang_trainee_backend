package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func TestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../test.db")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	if err := os.Remove("./../test.db"); err != nil {
		return err
	}
	return nil
}
