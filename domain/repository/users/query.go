package users

import (
	"github.com/bariasabda/test-flash-coffee/domain/entity"
)

func (r userRepository) FindByID(id int) (*entity.User, error) {
	var user entity.User
	err := r.db.Find(&user).Where("id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
