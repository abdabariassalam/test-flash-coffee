package recipe

import (
	"errors"

	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

func (s recipeSvc) CreateRecipe(input entity.CreateRecipe) (*entity.CreateRecipe, error) {

	// check if not exist create
	user, err := s.user.FindOrCreate(input.User.Name)
	if err != nil {
		return nil, err
	}
	// check if not exist create
	recipecategory, err := s.RecipeCategory.FindOrCreate(input.RecipeCategory.Name)
	if err != nil {
		return nil, err
	}
	// check recipe not duplicate
	recipeIds, err := s.recipe.FindByNameAndDescription(input.Name, input.Description)
	if err != nil {
		return nil, err
	}
	recipeCategoryRecipes, err := s.RecipeCategoryRecipe.FindDuplicate(*recipecategory.Id, *recipeIds)
	if err != nil {
		return nil, err
	}
	if len(*recipeCategoryRecipes) >= 1 {
		return nil, errors.New("recipe already exist")
	}
	recipe := &entity.Recipe{}
	err = s.agent.TxAgent(s.agent.GetDBMaster(), func(tx *gorm.DB) error {
		// create recipe
		recipe, err = s.recipe.Create(tx, entity.Recipe{
			Name:        input.Name,
			Description: input.Description,
			AuthorID:    &user.Id,
		})
		if err != nil {
			return err
		}

		// create recipe category recipe
		_, err = s.RecipeCategoryRecipe.Create(tx, entity.RecipeCategoryRecipe{
			RecipeCategoryID: *recipecategory.Id,
			RecipeID:         *recipe.Id,
		})

		return err
	})
	if err != nil {
		return nil, err
	}

	return &entity.CreateRecipe{
		Id:             recipe.Id,
		Name:           recipe.Name,
		Description:    recipe.Description,
		User:           *user,
		RecipeCategory: *recipecategory,
	}, nil
}
