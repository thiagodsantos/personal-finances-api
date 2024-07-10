package entity

import "github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"

type AccountType string

const (
	Bank       AccountType = "bank"
	CreditCard AccountType = "credit_card"
	DebitCard  AccountType = "debit_card"
)

type Account struct {
	ID          valueobjects.Id
	Name        valueobjects.Name
	Description valueobjects.Text
	Type        AccountType
}
