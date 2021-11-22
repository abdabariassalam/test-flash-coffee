package users

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
)

func (r userRepository) FindOrCreate(name string) (*entity.User, error) {
	user := entity.User{}

	if err := r.db.Where(entity.User{Name: name}).FirstOrCreate(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
