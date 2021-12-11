package stepingredient

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return repository{db}
}

type Repository interface {
	Create(tx *gorm.DB, input entity.StepIngredient) (*entity.StepIngredient, error)
}
