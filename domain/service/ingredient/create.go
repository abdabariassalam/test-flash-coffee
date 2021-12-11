package ingredient

import (
	"errors"

	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

func (s service) CreateIngredient(input entity.CreateIngredientRequest) (*entity.CreateIngredientResponse, error) {
	// check if not exist create
	ingredientCategory, err := s.ingredientCategory.FindOrCreate(input.IngredientCategory.Name, input.IngredientCategory.Description)
	if err != nil {
		return nil, err
	}
	ingredientIds, err := s.ingredient.FindByNameAndColor(input.Name, input.Color)
	if err != nil {
		return nil, err
	}
	ingredientCategoryIngredients, err := s.ingredientCategoryIngredient.FindDuplicate(*ingredientCategory.Id, *ingredientIds)
	if err != nil {
		return nil, err
	}
	if len(*ingredientCategoryIngredients) >= 1 {
		return nil, errors.New("ingredient already exist")
	}
	// upload s3
	err = s.aws.AddFileToS3(input.Image.ImageFile, input.Image.ImageName, input.Image.ImageContentType, input.Image.ImageSize)
	if err != nil {
		return nil, err
	}

	url := s.aws.GetUrl(input.Image.ImageName)

	ingredient := &entity.Ingredient{}
	err = s.agent.TxAgent(s.agent.GetDBMaster(), func(tx *gorm.DB) error {
		// create ingredient
		ingredient, err = s.ingredient.Create(tx, entity.Ingredient{
			Name:  input.Name,
			Color: input.Color,
			Img:   url,
		})
		if err != nil {
			return err
		}

		// create ingredient category ingredient
		_, err = s.ingredientCategoryIngredient.Create(tx, entity.IngredientCategoryIngredient{
			IngredientCategoryID: *ingredientCategory.Id,
			IngredientID:         *ingredient.Id,
		})

		return err
	})
	if err != nil {
		return nil, err
	}

	return &entity.CreateIngredientResponse{
		Id:    ingredient.Id,
		Name:  ingredient.Name,
		Color: ingredient.Color,
		IngredientCategory: entity.IngredientCategory{
			Id:          ingredientCategory.Id,
			Name:        ingredientCategory.Name,
			Description: ingredientCategory.Description,
		},
		Image: ingredient.Img,
	}, nil
}
