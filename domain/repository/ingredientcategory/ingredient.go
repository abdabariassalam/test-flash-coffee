package ingredientcategory

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
	FindOrCreate(name, description string) (*entity.IngredientCategory, error)
}
