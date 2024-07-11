package entity

import (
	"fmt"

	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type ExpenseRepeatTimesFrequency string

const (
	Monthly ExpenseRepeatTimesFrequency = "monthly"
	Yearly  ExpenseRepeatTimesFrequency = "yearly"
)

func (e ExpenseRepeatTimesFrequency) IsValid() error {
	switch e {
	case Monthly, Yearly:
		return nil
	}
	return fmt.Errorf("invalid expense repeat times frequency")
}

type Expense struct {
	Id                   valueobjects.Id
	AccountId            valueobjects.Id
	Amount               valueobjects.Money
	Date                 valueobjects.Date
	Description          valueobjects.Text
	Paid                 valueobjects.Flag
	Repeat               valueobjects.Flag
	RepeatTimes          valueobjects.Numeric
	RepeatTimesFrequency ExpenseRepeatTimesFrequency
}

func (e Expense) IsValid() error {
	var err []error

	err = append(err, e.Amount.IsValid())
	err = append(err, e.Date.IsValid())
	err = append(err, e.RepeatTimesFrequency.IsValid())

	for _, e := range err {
		if e != nil {
			return e
		}
	}

	return nil
}

func GetTotals(expenses []Expense) valueobjects.Money {
	var total valueobjects.Money
	for _, e := range expenses {
		total = total.Sum(e.Amount)
	}
	return total
}
