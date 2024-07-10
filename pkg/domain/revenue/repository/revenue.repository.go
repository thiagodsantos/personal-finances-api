package repository

import (
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/revenue/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type RevenueRepository interface {
	Create(revenue entity.Revenue) (entity.Revenue, error)
	DeleteById(id valueobjects.Id) error
	GetAll() ([]entity.Revenue, error)
	GetById(id valueobjects.Id) (entity.Revenue, error)
	Update(revenue entity.Revenue) (entity.Revenue, error)
}
