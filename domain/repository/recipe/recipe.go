package recipe

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

type recipeRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return recipeRepository{db}
}

type Repository interface {
	FindRecipes() (*[]entity.Recipes, error)
	Create(tx *gorm.DB, input entity.Recipe) (*entity.Recipe, error)
	FindByNameAndDescription(name, description string) (*[]int, error)
}
