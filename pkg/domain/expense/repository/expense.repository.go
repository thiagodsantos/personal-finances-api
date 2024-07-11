package repository

import (
	"fmt"

	"github.com/thiagodsantos/personal-finances-api/pkg/domain/expense/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type ExpenseRepository interface {
	Create(expense entity.Expense) error
	DeleteById(id valueobjects.Id) error
	GetAll() ([]entity.Expense, error)
	GetById(id valueobjects.Id) (entity.Expense, error)
	Update(expense entity.Expense) error
}

type repository struct{}

func (r repository) Create(expense entity.Expense) error {
	fmt.Println("ExpenseRepositoryImpl.Create")
	fmt.Println(expense)
	return nil
}

func (r repository) DeleteById(id valueobjects.Id) error {
	fmt.Println("ExpenseRepositoryImpl.DeleteById")
	fmt.Println(id)
	return nil
}

func (r repository) GetAll() ([]entity.Expense, error) {
	fmt.Println("ExpenseRepositoryImpl.GetAll")
	return []entity.Expense{}, nil
}

func (r repository) GetById(id valueobjects.Id) (entity.Expense, error) {
	fmt.Println("ExpenseRepositoryImpl.GetById")
	fmt.Println(id)
	return entity.Expense{}, nil
}

func (r repository) Update(expense entity.Expense) error {
	fmt.Println("ExpenseRepositoryImpl.Update")
	fmt.Println(expense)
	return nil
}

func NewExpenseRepository() ExpenseRepository {
	return repository{}
}
