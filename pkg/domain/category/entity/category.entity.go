package entity

import "github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"

type Category struct {
	Id            valueobjects.Id
	Description   valueobjects.Text
	Name          valueobjects.Name
	Subcategories []Subcategory
}

func (c Category) IsValid() error {
	err := c.Name.IsValid()
	if err != nil {
		return err
	}

	return nil
}
