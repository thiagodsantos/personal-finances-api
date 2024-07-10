package repository

import (
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/category/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type CategoryRepository interface {
	Create(category entity.Category) (entity.Category, error)
	DeleteById(id valueobjects.Id) error
	GetAll() ([]entity.Category, error)
	GetById(id valueobjects.Id) (entity.Category, error)
	Update(category entity.Category) (entity.Category, error)
	GetSubcategoriesByCategoryId(id valueobjects.Id) ([]entity.Subcategory, error)
}
