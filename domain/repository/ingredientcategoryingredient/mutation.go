package ingredientcategoryingredient

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

func (r repository) Create(tx *gorm.DB, input entity.IngredientCategoryIngredient) (*entity.IngredientCategoryIngredient, error) {
	err := tx.Create(&input).Error
	if err != nil {
		return nil, err
	}

	return &input, nil
}
