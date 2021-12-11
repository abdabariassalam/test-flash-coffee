package step

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

func (r repository) Create(tx *gorm.DB, input entity.Step) (*entity.Step, error) {
	db := tx.Create(&input)
	err := db.Error
	if err != nil {
		return nil, err
	}

	return &input, nil
}

func (r repository) UpdateStepNumber(tx *gorm.DB, id, recipeId, stepNumber int) error {
	db := tx.Exec("UPDATE step SET step_number = step_number + 1 where recipe_id = ? AND step_number >= ? and id <> ?", recipeId, stepNumber, id)
	err := db.Error
	return err
}
