package recipe

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/bariasabda/test-flash-coffee/domain/repository/recipe"
	"github.com/bariasabda/test-flash-coffee/domain/repository/users"
)

type recipeSvc struct {
	recipe recipe.Repository
	user   users.Repository
}

type RecipeSvc interface {
	GetRecipes() (*[]entity.Recipe, error)
}

func New(recipeRepo recipe.Repository, userRepo users.Repository) RecipeSvc {
	return recipeSvc{
		recipe: recipeRepo,
		user:   userRepo,
	}
}
