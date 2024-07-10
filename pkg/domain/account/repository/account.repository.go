package repository

import (
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/account/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type AccountRepository interface {
	Create(account entity.Account) (entity.Account, error)
	DeleteById(id valueobjects.Id) error
	GetAll() ([]entity.Account, error)
	GetById(id valueobjects.Id) (entity.Account, error)
	Update(account entity.Account) (entity.Account, error)
}
