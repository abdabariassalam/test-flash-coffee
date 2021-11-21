package users

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return userRepository{db}
}

type Repository interface {
	FindByID(id int) (*entity.User, error)
}
