package step

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/bariasabda/test-flash-coffee/domain/repository/agent"
	"github.com/bariasabda/test-flash-coffee/domain/repository/aws"
	"github.com/bariasabda/test-flash-coffee/domain/repository/ingredient"
	"github.com/bariasabda/test-flash-coffee/domain/repository/recipe"
	"github.com/bariasabda/test-flash-coffee/domain/repository/recipecategoryrecipe"
	"github.com/bariasabda/test-flash-coffee/domain/repository/step"
	"github.com/bariasabda/test-flash-coffee/domain/repository/stepingredient"
)

type service struct {
	agent                agent.Repository
	recipe               recipe.Repository
	ingredient           ingredient.Repository
	step                 step.Repository
	stepIngredient       stepingredient.Repository
	RecipeCategoryRecipe recipecategoryrecipe.Repository
	aws                  aws.Interface
}

type Service interface {
	CreateStep(input entity.CreateStep) (*entity.CreateStepResponse, error)
}

func New(
	agentRepo agent.Repository,
	recipeRepo recipe.Repository,
	ingredientRepo ingredient.Repository,
	stepRepo step.Repository,
	stepIngredientRepo stepingredient.Repository,
	recipeCategoryRecipeRepo recipecategoryrecipe.Repository,
	aws aws.Interface) Service {
	return service{
		agent:                agentRepo,
		recipe:               recipeRepo,
		ingredient:           ingredientRepo,
		step:                 stepRepo,
		stepIngredient:       stepIngredientRepo,
		RecipeCategoryRecipe: recipeCategoryRecipeRepo,
		aws:                  aws,
	}
}
