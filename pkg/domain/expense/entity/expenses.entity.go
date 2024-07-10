package entity

import "github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"

type ExpenseRepeatTimesFrequency string

const (
	Monthly ExpenseRepeatTimesFrequency = "monthly"
	Yearly  ExpenseRepeatTimesFrequency = "yearly"
)

func (e ExpenseRepeatTimesFrequency) IsValid() bool {
	switch e {
	case Monthly, Yearly:
		return true
	}
	return false
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

func GetTotals(expenses []Expense) valueobjects.Money {
	var total valueobjects.Money
	for _, e := range expenses {
		total = total.Sum(e.Amount)
	}
	return total
}
