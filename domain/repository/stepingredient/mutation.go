package stepingredient

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

func (r repository) Create(tx *gorm.DB, input entity.StepIngredient) (*entity.StepIngredient, error) {
	db := tx.Create(&input)
	err := db.Error
	if err != nil {
		return nil, err
	}

	return &input, nil
}
