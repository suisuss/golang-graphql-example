package models

import (
	"gorm.io/gorm"
	"github.com/twinj/uuid"
)

// We want our ids to be uuids, so we define that here

func (mod *Question) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	tx.Model(mod).Update("id", uuid.String())
	return
}

func (mod *QuestionOption) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	tx.Model(mod).Update("id", uuid.String())
	return
}

func (mod *Answer) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	tx.Model(mod).Update("id", uuid.String())
	return
}
