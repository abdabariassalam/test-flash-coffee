package ingredientcategory

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
)

func (r repository) FindOrCreate(name, description string) (*entity.IngredientCategory, error) {
	ingredientCategory := entity.IngredientCategory{}

	if err := r.db.Where(entity.IngredientCategory{Name: name, Description: description}).FirstOrCreate(&ingredientCategory).Error; err != nil {
		return nil, err
	}

	return &ingredientCategory, nil
}
