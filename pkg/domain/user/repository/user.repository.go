package repository

import (
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/user/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	GetByEmail(email valueobjects.Email) (entity.User, error)
	Update(user entity.User) (entity.User, error)
}
