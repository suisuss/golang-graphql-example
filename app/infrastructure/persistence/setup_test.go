package persistence_test

import (
	"fmt"
	"gorm.io/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"github.com/suisuss/golang-graphql-example/app/models"
	"os"
	"gorm.io/driver/sqlite"
)

func DBConn() (*gorm.DB, error) {

	if _, err := os.Stat("./../../../.env"); !os.IsNotExist(err) {
		var err error
		err = godotenv.Load(os.ExpandEnv("./../../../.env"))
		if err != nil {
			log.Fatalf("Error getting env %v\n", err)
		} else {
			fmt.Println("we have the env")
		}
	}
	db, err := LocalDatabase()
	return db, err
}
//Local DB
func LocalDatabase() (*gorm.DB, error) {

	test_database := os.Getenv("TEST_DATABASE")

	db, err := gorm.Open(sqlite.Open(test_database), &gorm.Config{})
	if err != nil {
		return nil, err
	} else {
		log.Println("CONNECTED TO: ", test_database)
	}

	db.Migrator().DropTable(&models.Question{}, &models.Answer{}, &models.QuestionOption{})

	db.Debug()
	db.AutoMigrate(
		models.Question{},
		models.Answer{},
		models.QuestionOption{},
	)
	return db, nil
}

func seedQuestion(db *gorm.DB) (*models.Question, error) {
	question := &models.Question{
		ID:    "1",
		Title: "First Question",
	}
	err := db.Create(&question).Error
	if err != nil {
		return nil, err
	}
	return question, nil
}

func seedQuestions(db *gorm.DB) ([]models.Question, error) {
	questions := []models.Question{
		{
			ID:    "1",
			Title: "First Question",
		},
		{
			ID:    "2",
			Title: "Second Question",
		},
	}
	for _, v := range questions {
		err := db.Create(&v).Error
		if err != nil {
			return nil, err
		}
	}
	return questions, nil
}

func seedQuestionOption(db *gorm.DB) (*models.QuestionOption, error) {
	quesOpt := &models.QuestionOption{
		ID:         "1",
		QuestionID: "1",
		Title:      "Option 1",
		Position:   1,
		IsCorrect:  false,
	}
	err := db.Create(&quesOpt).Error
	if err != nil {
		return nil, err
	}
	return quesOpt, nil
}

func seedQuestionOptions(db *gorm.DB) ([]models.QuestionOption, error) {
	quesOpts := []models.QuestionOption{
		{
			ID:         "1",
			QuestionID: "1",
			Title:      "Option 1",
			Position:   1,
			IsCorrect:  false,
		},
		{
			ID:         "2",
			QuestionID: "2",
			Title:      "Option 2",
			Position:   2,
			IsCorrect:  true,
		},
	}
	for _, v := range quesOpts {
		err := db.Create(&v).Error
		if err != nil {
			return nil, err
		}
	}
	return quesOpts, nil
}

func seedAnswer(db *gorm.DB) (*models.Answer, error) {
	ans := &models.Answer{
		QuestionID: "1",
		OptionID:   "1",
		IsCorrect:  true,
	}
	err := db.Create(&ans).Error
	if err != nil {
		return nil, err
	}
	return ans, nil
}

func seedAnswers(db *gorm.DB) ([]models.Answer, error) {
	answers := []models.Answer{
		{
			QuestionID: "1",
			OptionID:   "1",
			IsCorrect:  false,
		},
		{
			QuestionID: "1",
			OptionID:   "2",
			IsCorrect:  true,
		},
	}
	for _, v := range answers {
		err := db.Create(&v).Error
		if err != nil {
			return nil, err
		}
	}

	return answers, nil
}
