package ingredient

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

type ingredientRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return ingredientRepository{db}
}

type Repository interface {
	Create(tx *gorm.DB, input entity.Ingredient) (*entity.Ingredient, error)
	FindByNameAndColor(name string, color int) (*[]int, error)
	FindIngredients() (*[]entity.Ingredient, error)
	FindIngredientByID(id int) (*entity.Ingredient, error)
}
