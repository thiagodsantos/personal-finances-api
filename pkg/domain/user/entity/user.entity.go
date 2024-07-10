package entity

import "github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"

type User struct {
	Id    valueobjects.Id
	Name  valueobjects.Name
	Email valueobjects.Email
}
