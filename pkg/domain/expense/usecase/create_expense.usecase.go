package usecase

import (
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/expense/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/expense/repository"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type CreateExpenseInput struct {
	Amount               valueobjects.Money
	Date                 valueobjects.Date
	Description          valueobjects.Text
	Paid                 valueobjects.Flag
	Repeat               valueobjects.Flag
	RepeatTimes          valueobjects.Numeric
	RepeatTimesFrequency entity.ExpenseRepeatTimesFrequency
}

type CreateExpense struct {
	repo repository.ExpenseRepository
}

func (ce *CreateExpense) Execute(input CreateExpenseInput) error {
	expense := entity.Expense{
		Amount:               input.Amount,
		Date:                 input.Date,
		Description:          input.Description,
		Paid:                 input.Paid,
		Repeat:               input.Repeat,
		RepeatTimes:          input.RepeatTimes,
		RepeatTimesFrequency: input.RepeatTimesFrequency,
	}

	err := expense.IsValid()
	if err != nil {
		return err
	}

	err = ce.repo.Create(expense)
	if err != nil {
		return err
	}

	return nil
}
