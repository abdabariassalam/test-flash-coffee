package ingredientcategoryingredient

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
)

func (r repository) FindDuplicate(IngredientCategoryId int, ingredientIds []int) (*[]entity.IngredientCategoryIngredient, error) {
	ingredientCategoryIngredient := []entity.IngredientCategoryIngredient{}

	if err := r.db.Where("ingredient_category_id = ? and ingredient_id in (?)", IngredientCategoryId, ingredientIds).Find(&ingredientCategoryIngredient).Error; err != nil {
		return nil, err
	}

	return &ingredientCategoryIngredient, nil
}
