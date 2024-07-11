package usecase

import (
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/revenue/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/revenue/repository"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type CreateRevenueInput struct {
	Amount               valueobjects.Money
	Date                 valueobjects.Date
	Description          valueobjects.Text
	Received             valueobjects.Flag
	Repeat               valueobjects.Flag
	RepeatTimes          valueobjects.Numeric
	RepeatTimesFrequency entity.RevenueRepeatTimesFrequency
}

type CreateRevenue struct {
	repo repository.RevenueRepository
}

func (cr *CreateRevenue) Execute(input CreateRevenueInput) error {
	revenue := entity.Revenue{
		Amount:               input.Amount,
		Date:                 input.Date,
		Description:          input.Description,
		Received:             input.Received,
		Repeat:               input.Repeat,
		RepeatTimes:          input.RepeatTimes,
		RepeatTimesFrequency: input.RepeatTimesFrequency,
	}

	err := revenue.IsValid()
	if err != nil {
		return err
	}

	err = cr.repo.Create(revenue)
	if err != nil {
		return err
	}

	return nil
}
