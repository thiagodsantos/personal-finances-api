package repository

import (
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/expense/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type ExpenseRepository interface {
	Create(expense entity.Expense) (entity.Expense, error)
	DeleteById(id valueobjects.Id) error
	GetAll() ([]entity.Expense, error)
	GetById(id valueobjects.Id) (entity.Expense, error)
	Update(expense entity.Expense) (entity.Expense, error)
}
