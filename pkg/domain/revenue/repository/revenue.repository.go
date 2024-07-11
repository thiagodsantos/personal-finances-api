package repository

import (
	"fmt"

	"github.com/thiagodsantos/personal-finances-api/pkg/domain/revenue/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type RevenueRepository interface {
	Create(revenue entity.Revenue) error
	DeleteById(id valueobjects.Id) error
	GetAll() ([]entity.Revenue, error)
	GetById(id valueobjects.Id) (entity.Revenue, error)
	Update(revenue entity.Revenue) error
}

type repository struct{}

func (r repository) Create(revenue entity.Revenue) error {
	fmt.Println("RevenueRepositoryImpl.Create")
	fmt.Println(revenue)
	return nil
}

func (r repository) DeleteById(id valueobjects.Id) error {
	fmt.Println("RevenueRepositoryImpl.DeleteById")
	fmt.Println(id)
	return nil
}

func (r repository) GetAll() ([]entity.Revenue, error) {
	fmt.Println("RevenueRepositoryImpl.GetAll")
	return []entity.Revenue{}, nil
}

func (r repository) GetById(id valueobjects.Id) (entity.Revenue, error) {
	fmt.Println("RevenueRepositoryImpl.GetById")
	fmt.Println(id)
	return entity.Revenue{}, nil
}

func (r repository) Update(revenue entity.Revenue) error {
	fmt.Println("RevenueRepositoryImpl.Update")
	fmt.Println(revenue)
	return nil
}

func NewRevenueRepository() RevenueRepository {
	return repository{}
}
