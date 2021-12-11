package step

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
	FindByRecipeID(recipeId int) (res *[]entity.Step, totalRows int, err error)
	Create(tx *gorm.DB, input entity.Step) (*entity.Step, error)
	UpdateStepNumber(tx *gorm.DB, id, recipeId, stepNumber int) error
}
