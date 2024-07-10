package entity

import "github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"

type RevenueRepeatTimesFrequency string

const (
	Monthly RevenueRepeatTimesFrequency = "monthly"
	Yearly  RevenueRepeatTimesFrequency = "yearly"
)

func (r RevenueRepeatTimesFrequency) IsValid() bool {
	switch r {
	case Monthly, Yearly:
		return true
	}
	return false
}

type Revenue struct {
	Id                   valueobjects.Id
	AccountId            valueobjects.Id
	Amount               valueobjects.Money
	Date                 valueobjects.Date
	Description          valueobjects.Text
	Received             valueobjects.Flag
	Repeat               valueobjects.Flag
	RepeatTimes          valueobjects.Numeric
	RepeatTimesFrequency RevenueRepeatTimesFrequency
}

func GetTotals(revenues []Revenue) valueobjects.Money {
	var total valueobjects.Money
	for _, r := range revenues {
		total = total.Sum(r.Amount)
	}
	return total
}
