package repository

import (
	"fmt"

	"github.com/thiagodsantos/personal-finances-api/pkg/domain/user/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type UserRepository interface {
	Create(user entity.User) error
	GetByEmail(email valueobjects.Email) (entity.User, error)
	Update(user entity.User) error
}

type repository struct{}

func (r repository) Create(user entity.User) error {
	fmt.Println("UserRepositoryImpl.Create")
	fmt.Println(user)
	return nil
}

func (r repository) GetByEmail(email valueobjects.Email) (entity.User, error) {
	fmt.Println("UserRepositoryImpl.GetByEmail")
	fmt.Println(email)
	return entity.User{}, nil
}

func (r repository) Update(user entity.User) error {
	fmt.Println("UserRepositoryImpl.Update")
	fmt.Println(user)
	return nil
}

func NewUserRepository() UserRepository {
	return repository{}
}
