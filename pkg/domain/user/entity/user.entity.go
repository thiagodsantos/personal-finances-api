package entity

import "github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"

type User struct {
	Id    valueobjects.Id
	Name  valueobjects.Name
	Email valueobjects.Email
}

func (u User) IsValid() error {
	var err []error

	err = append(err, u.Name.IsValid())
	err = append(err, u.Email.IsValid())

	for _, e := range err {
		if e != nil {
			return e
		}
	}

	return nil
}
