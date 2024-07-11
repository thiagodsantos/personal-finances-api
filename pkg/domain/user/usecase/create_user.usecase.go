package usecase

import (
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/user/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/user/repository"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type CreateUserInput struct {
	Name  valueobjects.Name
	Email valueobjects.Email
}

type CreateUser struct {
	repo repository.UserRepository
}

func (cu *CreateUser) Execute(input CreateUserInput) error {
	user := entity.User{
		Name:  input.Name,
		Email: input.Email,
	}

	err := user.IsValid()
	if err != nil {
		return err
	}

	err = cu.repo.Create(user)
	if err != nil {
		return err
	}

	return nil
}
