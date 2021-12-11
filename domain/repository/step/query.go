package step

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
)

func (r repository) FindByRecipeID(recipeId int) (res *[]entity.Step, totalRows int, err error) {
	var steps []entity.Step
	respDB := r.db.Model(steps).Where("recipe_id = ?", recipeId)
	respDB.Find(&steps)
	if respDB.Error != nil {
		return nil, 0, respDB.Error
	}
	totalRows = len(steps)
	return &steps, totalRows, nil
}
