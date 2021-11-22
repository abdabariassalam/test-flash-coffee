package recipecategory

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
)

func (r recipeCategoryRepository) FindOrCreate(name string) (*entity.RecipeCategory, error) {
	recipeCategory := entity.RecipeCategory{}

	if err := r.db.Where(entity.RecipeCategory{Name: name}).FirstOrCreate(&recipeCategory).Error; err != nil {
		return nil, err
	}

	return &recipeCategory, nil
}
