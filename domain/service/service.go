package service

import (
	"github.com/bariasabda/test-flash-coffee/domain/repository"
	"github.com/bariasabda/test-flash-coffee/domain/service/ingredient"
	"github.com/bariasabda/test-flash-coffee/domain/service/recipe"
	"github.com/bariasabda/test-flash-coffee/domain/service/step"
)

type Service struct {
	Recipe     recipe.RecipeSvc
	Ingredient ingredient.Service
	Step       step.Service
}

func New(repo repository.Repository) Service {
	return Service{
		Recipe:     recipe.New(repo.Agent, repo.Recipe, repo.User, repo.RecipeCategory, repo.RecipeCategoryRecipe),
		Ingredient: ingredient.New(repo.AWS, repo.Ingredient, repo.IngredientCategory, repo.IngredientCategoryIngredient, repo.Agent),
		Step:       step.New(repo.Agent, repo.Recipe, repo.Ingredient, repo.Step, repo.StepIngredient, repo.RecipeCategoryRecipe, repo.AWS),
	}
}
