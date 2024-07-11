package repository

import (
	"fmt"

	"github.com/thiagodsantos/personal-finances-api/pkg/domain/account/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type AccountRepository interface {
	Create(account entity.Account) error
	DeleteById(id valueobjects.Id) error
	GetAll() ([]entity.Account, error)
	GetById(id valueobjects.Id) (entity.Account, error)
	Update(account entity.Account) error
}

type repository struct{}

func (r repository) Create(account entity.Account) error {
	fmt.Println("AccountRepositoryImpl.Create")
	fmt.Println(account)
	return nil
}

func (r repository) DeleteById(id valueobjects.Id) error {
	fmt.Println("AccountRepositoryImpl.DeleteById")
	fmt.Println(id)
	return nil
}

func (r repository) GetAll() ([]entity.Account, error) {
	fmt.Println("AccountRepositoryImpl.GetAll")
	return []entity.Account{}, nil
}

func (r repository) GetById(id valueobjects.Id) (entity.Account, error) {
	fmt.Println("AccountRepositoryImpl.GetById")
	fmt.Println(id)
	return entity.Account{}, nil
}

func (r repository) Update(account entity.Account) error {
	fmt.Println("AccountRepositoryImpl.Update")
	fmt.Println(account)
	return nil
}

func NewAccountRepository() AccountRepository {
	return repository{}
}
