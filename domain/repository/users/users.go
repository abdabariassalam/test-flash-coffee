package users

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return userRepository{db}
}

type Repository interface {
	FindOrCreate(name string) (*entity.User, error)
}
