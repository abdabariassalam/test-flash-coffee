package repository

import (
	"github.com/bariasabda/test-flash-coffee/domain/repository/recipe"
	"github.com/bariasabda/test-flash-coffee/domain/repository/users"
	"gorm.io/gorm"
)

type Repository struct {
	Recipe recipe.Repository
	User   users.Repository
}

func New(db *gorm.DB) Repository {
	return Repository{
		Recipe: recipe.New(db),
		User:   users.New(db),
	}
}
