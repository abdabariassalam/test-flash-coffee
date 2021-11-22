package recipe

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
)

func (r recipeRepository) FindRecipes() (*[]entity.Recipes, error) {
	recipes := []entity.Recipes{}
	rows, err := r.db.Table("recipe").
		Select("recipe.id, recipe.name, recipe.description, users.id, users.name").
		Joins("Join users on users.id = recipe.author_id").Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		recipe := entity.Recipes{}
		err = rows.Scan(
			&recipe.Id,
			&recipe.Name,
			&recipe.Description,
			&recipe.Author.Id,
			&recipe.Author.Name)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	return &recipes, nil
}

func (r recipeRepository) FindByNameAndDescription(name, description string) (*[]int, error) {
	var Ids []int
	rows, err := r.db.Table("recipe").
		Select("id").Where("name = ? and description = ?", name, description).Rows()
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
