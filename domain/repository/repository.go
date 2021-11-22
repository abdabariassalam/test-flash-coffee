package repository

import (
	"github.com/bariasabda/test-flash-coffee/domain/repository/agent"
	"github.com/bariasabda/test-flash-coffee/domain/repository/recipe"
	"github.com/bariasabda/test-flash-coffee/domain/repository/recipecategory"
	"github.com/bariasabda/test-flash-coffee/domain/repository/recipecategoryrecipe"
	"github.com/bariasabda/test-flash-coffee/domain/repository/users"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	Agent                agent.Repository
	Recipe               recipe.Repository
	User                 users.Repository
	RecipeCategory       recipecategory.Repository
	RecipeCategoryRecipe recipecategoryrecipe.Repository
}

func New(db *gorm.DB) Repository {
	return Repository{
		Recipe:               recipe.New(db),
		User:                 users.New(db),
		RecipeCategory:       recipecategory.New(db),
		RecipeCategoryRecipe: recipecategoryrecipe.New(db),
		Agent:                agent.New(db),
	}
}
