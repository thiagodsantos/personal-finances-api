package category

import "github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"

type Category struct {
	ID            valueobjects.Id
	Description   valueobjects.Text
	Name          valueobjects.Name
	Subcategories []Category
}
