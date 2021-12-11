package ingredient

import "github.com/bariasabda/test-flash-coffee/domain/entity"

func (s service) GetIngredients() (*[]entity.Ingredient, error) {
	ingredients, err := s.ingredient.FindIngredients()
	return ingredients, err
}
