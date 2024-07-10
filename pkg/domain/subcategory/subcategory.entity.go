package subcategory

import "github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"

type Subcategory struct {
	Id          valueobjects.Id
	CategoryId  valueobjects.Id
	Description valueobjects.Text
	Name        valueobjects.Name
}
