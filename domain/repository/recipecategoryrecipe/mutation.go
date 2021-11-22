package recipecategoryrecipe

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

func (r recipeCategoryRecipeRepository) Create(tx *gorm.DB, input entity.RecipeCategoryRecipe) (*entity.RecipeCategoryRecipe, error) {
	err := tx.Create(&input).Error
	if err != nil {
		return nil, err
	}

	return &input, nil
}
