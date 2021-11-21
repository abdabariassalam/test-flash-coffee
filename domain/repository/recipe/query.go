package recipe

import (
	"log"

	"github.com/bariasabda/test-flash-coffee/domain/entity"
)

func (r recipeRepository) FindRecipes() (*[]entity.Recipe, error) {
	recipes := []entity.Recipe{}
	rows, err := r.db.Table("recipe").
		Select("recipe.id, recipe.name, recipe.description, users.id, users.name").
		Joins("Join users on users.id = recipe.author_id").Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		recipe := entity.Recipe{}
		err = rows.Scan(
			&recipe.Id,
			&recipe.Name,
			&recipe.Description,
			&recipe.Author.Id,
			&recipe.Author.Name)
		if err != nil {
			log.Panic(err)
		}
		recipes = append(recipes, recipe)
	}
	return &recipes, nil
}
