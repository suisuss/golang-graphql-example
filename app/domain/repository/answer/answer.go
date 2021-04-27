package answer

import (
	"github.com/suisuss/golang-graphql-example/app/models"
)

type AnsService interface {
	CreateAnswer(answer *models.Answer) (*models.Answer, error)
	UpdateAnswer(answer *models.Answer) (*models.Answer, error)
	DeleteAnswer(id string) error
	GetAnswerByID(id string) (*models.Answer, error)
	GetAllQuestionAnswers(questionId string) ([]*models.Answer, error)
}

