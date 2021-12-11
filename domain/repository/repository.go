package repository

import (
	"log"

	awspkg "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bariasabda/test-flash-coffee/config"
	"github.com/bariasabda/test-flash-coffee/domain/repository/agent"
	"github.com/bariasabda/test-flash-coffee/domain/repository/aws"
	"github.com/bariasabda/test-flash-coffee/domain/repository/ingredient"
	"github.com/bariasabda/test-flash-coffee/domain/repository/ingredientcategory"
	"github.com/bariasabda/test-flash-coffee/domain/repository/ingredientcategoryingredient"
	"github.com/bariasabda/test-flash-coffee/domain/repository/recipe"
	"github.com/bariasabda/test-flash-coffee/domain/repository/recipecategory"
	"github.com/bariasabda/test-flash-coffee/domain/repository/recipecategoryrecipe"
	"github.com/bariasabda/test-flash-coffee/domain/repository/step"
	"github.com/bariasabda/test-flash-coffee/domain/repository/stepingredient"
	"github.com/bariasabda/test-flash-coffee/domain/repository/users"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	Agent                        agent.Repository
	Recipe                       recipe.Repository
	User                         users.Repository
	RecipeCategory               recipecategory.Repository
	RecipeCategoryRecipe         recipecategoryrecipe.Repository
	Ingredient                   ingredient.Repository
	IngredientCategory           ingredientcategory.Repository
	IngredientCategoryIngredient ingredientcategoryingredient.Repository
	AWS                          aws.Interface
	Step                         step.Repository
	StepIngredient               stepingredient.Repository
}

type Args struct {
	DB  *gorm.DB
	Cfg *config.Config
}

func New(a *Args) Repository {
	s, err := session.NewSession(&awspkg.Config{
		Region: awspkg.String(a.Cfg.AWS.Region),
		Credentials: credentials.NewStaticCredentials(
			a.Cfg.AWS.AccessKeyID,
			a.Cfg.AWS.SecretAccessKey,
			"",
		),
	})

	if err != nil {
		log.Fatal(err)
	}
	return Repository{
		Recipe:                       recipe.New(a.DB),
		User:                         users.New(a.DB),
		RecipeCategory:               recipecategory.New(a.DB),
		RecipeCategoryRecipe:         recipecategoryrecipe.New(a.DB),
		Agent:                        agent.New(a.DB),
		AWS:                          aws.New(a.Cfg, s),
		Ingredient:                   ingredient.New(a.DB),
		IngredientCategory:           ingredientcategory.New(a.DB),
		IngredientCategoryIngredient: ingredientcategoryingredient.New(a.DB),
		Step:                         step.New(a.DB),
		StepIngredient:               stepingredient.New(a.DB),
	}
}
