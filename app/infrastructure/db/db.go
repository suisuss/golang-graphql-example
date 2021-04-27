package db

import (
	"gorm.io/driver/sqlite"
  "gorm.io/gorm"
	"log"
	"github.com/suisuss/golang-graphql-example/app/models"
)

func OpenDB(database string) *gorm.DB {

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		log.Fatalf("%s", err)
	}

	db.AutoMigrate(&models.Question{}, &models.QuestionOption{}, &models.Answer{})

	return db
}