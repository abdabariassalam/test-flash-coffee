package ingredient

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/bariasabda/test-flash-coffee/domain/repository/agent"
	"github.com/bariasabda/test-flash-coffee/domain/repository/aws"
	"github.com/bariasabda/test-flash-coffee/domain/repository/ingredient"
	"github.com/bariasabda/test-flash-coffee/domain/repository/ingredientcategory"
	"github.com/bariasabda/test-flash-coffee/domain/repository/ingredientcategoryingredient"
)

type service struct {
	aws                          aws.Interface
	ingredient                   ingredient.Repository
	ingredientCategory           ingredientcategory.Repository
	ingredientCategoryIngredient ingredientcategoryingredient.Repository
	agent                        agent.Repository
}

type Service interface {
	CreateIngredient(input entity.CreateIngredientRequest) (*entity.CreateIngredientResponse, error)
	GetIngredients() (*[]entity.Ingredient, error)
}

func New(
	awsInterface aws.Interface,
	ingredientRepo ingredient.Repository,
	ingredientCategoryRepo ingredientcategory.Repository,
	ingredientCategoryIngredientRepo ingredientcategoryingredient.Repository,
	agentRepo agent.Repository) service {
	return service{
		aws:                          awsInterface,
		ingredient:                   ingredientRepo,
		ingredientCategory:           ingredientCategoryRepo,
		ingredientCategoryIngredient: ingredientCategoryIngredientRepo,
		agent:                        agentRepo,
	}
}
