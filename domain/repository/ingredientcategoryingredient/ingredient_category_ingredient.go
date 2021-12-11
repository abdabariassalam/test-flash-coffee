package ingredientcategoryingredient

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
	FindDuplicate(IngredientCategoryId int, ingredientIds []int) (*[]entity.IngredientCategoryIngredient, error)
	Create(tx *gorm.DB, input entity.IngredientCategoryIngredient) (*entity.IngredientCategoryIngredient, error)
}
