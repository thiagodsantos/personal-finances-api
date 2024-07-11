package entity

import (
	"fmt"

	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type AccountType string

const (
	Bank       AccountType = "bank"
	CreditCard AccountType = "credit_card"
	DebitCard  AccountType = "debit_card"
)

func (a AccountType) IsValid() error {
	switch a {
	case Bank, CreditCard, DebitCard:
		return nil
	default:
		return fmt.Errorf("invalid account type")
	}
}

type Account struct {
	Id          valueobjects.Id
	Name        valueobjects.Name
	Description valueobjects.Text
	Type        AccountType
}

func (a Account) IsValid() error {
	var err []error

	err = append(err, a.Name.IsValid())
	err = append(err, a.Type.IsValid())

	for _, e := range err {
		if e != nil {
			return e
		}
	}

	return nil
}
