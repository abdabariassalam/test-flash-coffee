package recipecategoryrecipe

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
)

func (r recipeCategoryRecipeRepository) FindDuplicate(recipeCategory int, recipeIds []int) (*[]entity.RecipeCategoryRecipe, error) {
	recipeCategoryRecipe := []entity.RecipeCategoryRecipe{}

	if err := r.db.Where("recipe_category_id = ? and recipe_id in (?)", recipeCategory, recipeIds).Find(&recipeCategoryRecipe).Error; err != nil {
		return nil, err
	}

	return &recipeCategoryRecipe, nil
}
