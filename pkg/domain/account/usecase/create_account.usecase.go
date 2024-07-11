package usecase

import (
	"fmt"

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

	exists, err := ca.repo.GetByName(account.Name, account.Type)
	if err != nil {
		return err
	}

	if exists.Id != "" {
		return fmt.Errorf("account already exists with name %s with type %s", account.Name, account.Type)
	}

	err = ca.repo.Create(account)
	if err != nil {
		return err
	}

	return nil
}
