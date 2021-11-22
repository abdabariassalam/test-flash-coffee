package recipe

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/bariasabda/test-flash-coffee/domain/repository/agent"
	"github.com/bariasabda/test-flash-coffee/domain/repository/recipe"
	"github.com/bariasabda/test-flash-coffee/domain/repository/recipecategory"
	"github.com/bariasabda/test-flash-coffee/domain/repository/recipecategoryrecipe"
	"github.com/bariasabda/test-flash-coffee/domain/repository/users"
)

type recipeSvc struct {
	agent                agent.Repository
	recipe               recipe.Repository
	user                 users.Repository
	RecipeCategory       recipecategory.Repository
	RecipeCategoryRecipe recipecategoryrecipe.Repository
}

type RecipeSvc interface {
	GetRecipes() (*[]entity.Recipes, error)
	CreateRecipe(input entity.CreateRecipe) (*entity.CreateRecipe, error)
}

func New(
	agentRepo agent.Repository,
	recipeRepo recipe.Repository,
	userRepo users.Repository,
	recipeCategoryRepo recipecategory.Repository,
	recipeCategoryRecipeRepo recipecategoryrecipe.Repository) RecipeSvc {
	return recipeSvc{
		agent:                agentRepo,
		recipe:               recipeRepo,
		user:                 userRepo,
		RecipeCategory:       recipeCategoryRepo,
		RecipeCategoryRecipe: recipeCategoryRecipeRepo,
	}
}
