package recipe

import "github.com/bariasabda/test-flash-coffee/domain/entity"

func (s recipeSvc) GetRecipes() (*[]entity.Recipe, error) {
	recipes, err := s.recipe.FindRecipes()
	return recipes, err
}
