package recipe

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

func (r recipeRepository) Create(tx *gorm.DB, input entity.Recipe) (*entity.Recipe, error) {
	db := tx.Create(&input)
	err := db.Error
	if err != nil {
		return nil, err
	}

	return &input, nil
}
