package ingredient

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
)

func (r ingredientRepository) FindByNameAndColor(name string, color int) (*[]int, error) {
	var Ids []int
	rows, err := r.db.Table("ingredient").
		Select("id").Where("name = ? and color = ?", name, color).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		recipe := entity.Recipes{}
		err = rows.Scan(&recipe.Id)
		if err != nil {
			return nil, err
		}
		Ids = append(Ids, recipe.Id)
	}
	return &Ids, nil
}

func (r ingredientRepository) FindIngredients() (*[]entity.Ingredient, error) {
	ingredients := []entity.Ingredient{}
	err := r.db.Find(&ingredients).Error
	if err != nil {
		return nil, err
	}
	return &ingredients, nil
}

func (r ingredientRepository) FindIngredientByID(id int) (*entity.Ingredient, error) {
	ingredients := entity.Ingredient{}
	err := r.db.First(&ingredients, id).Error
	if err != nil {
		return nil, err
	}
	return &ingredients, nil
}
