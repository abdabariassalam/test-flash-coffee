package recipe

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"gorm.io/gorm"
)

type recipeRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return recipeRepository{db}
}

type Repository interface {
	FindRecipes() (*[]entity.Recipe, error)
}
