package step

import (
	"fmt"

	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

func (s service) CreateStep(input entity.CreateStep) (*entity.CreateStepResponse, error) {
	// check recipe
	recipe, err := s.recipe.FindByID(input.StepIngredient.RecipeId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Get Recipe Error Reason: %v", err.Error())
		}
		return nil, err
	}
	// check ingredient
	ingredient, err := s.ingredient.FindIngredientByID(input.StepIngredient.IngredientId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Get Ingredient Error Reason: %v", err.Error())
		}
		return nil, err
	}
	// check totalRow untuk menentukan step_number
	_, totalRow, err := s.step.FindByRecipeID(input.StepIngredient.RecipeId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Get Ingredient Error Reason: %v", err.Error())
		}
		return nil, err
	}
	if input.StepNumber > totalRow+1 {
		return nil, fmt.Errorf("step number out of range")
	}
	// upload s3
	err = s.aws.AddFileToS3(input.Image.ImageFile, input.Image.ImageName, input.Image.ImageContentType, input.Image.ImageSize)
	if err != nil {
		return nil, err
	}
	url := s.aws.GetUrl(input.Image.ImageName)
	step := &entity.Step{}
	err = s.agent.TxAgent(s.agent.GetDBMaster(), func(tx *gorm.DB) error {
		// create recipe
		step, err = s.step.Create(tx, entity.Step{
			RecipeId:    input.StepIngredient.RecipeId,
			StepNumber:  input.StepNumber,
			Description: input.Description,
			Timer:       input.Timer,
			Image:       url,
		})
		if err != nil {
			return err
		}
		if step.StepNumber <= totalRow {
			err := s.step.UpdateStepNumber(tx, step.Id, step.RecipeId, step.StepNumber)
			if err != nil {
				return err
			}
		}

		// create recipe category recipe
		_, err = s.stepIngredient.Create(tx, entity.StepIngredient{
			RecipeId:     recipe.Id,
			IngredientId: *ingredient.Id,
			StepId:       step.Id,
			Amount:       input.StepIngredient.Amount,
			Unit:         input.StepIngredient.Unit,
		})

		return err
	})
	if err != nil {
		return nil, err
	}

	return &entity.CreateStepResponse{
		Id:          &step.Id,
		RecipeId:    step.RecipeId,
		StepNumber:  step.StepNumber,
		Description: step.Description,
		Timer:       step.Timer,
		StepIngredient: entity.StepIngredient{
			RecipeId:     recipe.Id,
			IngredientId: *ingredient.Id,
			StepId:       step.Id,
			Amount:       input.StepIngredient.Amount,
			Unit:         input.StepIngredient.Unit,
		},
		Image: url,
	}, nil
}
