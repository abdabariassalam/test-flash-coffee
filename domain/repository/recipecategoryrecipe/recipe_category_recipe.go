package recipecategoryrecipe

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

type recipeCategoryRecipeRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return recipeCategoryRecipeRepository{db}
}

type Repository interface {
	FindDuplicate(recipeCategory int, recipeIds []int) (*[]entity.RecipeCategoryRecipe, error)
	Create(tx *gorm.DB, input entity.RecipeCategoryRecipe) (*entity.RecipeCategoryRecipe, error)
}
