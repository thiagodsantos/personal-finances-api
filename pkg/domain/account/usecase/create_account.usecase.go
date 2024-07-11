package usecase

import (
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/account/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/account/repository"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type CreateAccountInput struct {
	Name        valueobjects.Name
	Description valueobjects.Text
	Type        entity.AccountType
}

type CreateAccount struct {
	repo repository.AccountRepository
}

func (ca *CreateAccount) Execute(input CreateAccountInput) error {
	account := entity.Account{
		Name:        input.Name,
		Description: input.Description,
		Type:        input.Type,
	}

	err := account.IsValid()
	if err != nil {
		return err
	}

	err = ca.repo.Create(account)
	if err != nil {
		return err
	}

	return nil
}
