package ingredient

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

func (r ingredientRepository) Create(tx *gorm.DB, input entity.Ingredient) (*entity.Ingredient, error) {
	db := tx.Create(&input)
	err := db.Error
	if err != nil {
		return nil, err
	}

	return &input, nil
}
