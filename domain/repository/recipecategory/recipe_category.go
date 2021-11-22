package recipecategory

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

type recipeCategoryRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return recipeCategoryRepository{db}
}

type Repository interface {
	FindOrCreate(name string) (*entity.RecipeCategory, error)
}
