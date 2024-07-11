package entity

import (
	"fmt"

	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type RevenueRepeatTimesFrequency string

const (
	Monthly RevenueRepeatTimesFrequency = "monthly"
	Yearly  RevenueRepeatTimesFrequency = "yearly"
)

func (r RevenueRepeatTimesFrequency) IsValid() error {
	switch r {
	case Monthly, Yearly:
		return nil
	}
	return fmt.Errorf("invalid revenue repeat times frequency")
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

func (e Revenue) IsValid() error {
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

func GetTotals(revenues []Revenue) valueobjects.Money {
	var total valueobjects.Money
	for _, r := range revenues {
		total = total.Sum(r.Amount)
	}
	return total
}
