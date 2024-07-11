package repository

import (
	"fmt"

	"github.com/thiagodsantos/personal-finances-api/pkg/domain/category/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type CategoryRepository interface {
	Create(category entity.Category) error
	DeleteById(id valueobjects.Id) error
	GetAll() ([]entity.Category, error)
	GetById(id valueobjects.Id) (entity.Category, error)
	Update(category entity.Category) error
	GetSubcategoriesByCategoryId(id valueobjects.Id) ([]entity.Subcategory, error)
}

type repository struct{}

func (r repository) Create(category entity.Category) error {
	fmt.Println("CategoryRepositoryImpl.Create")
	fmt.Println(category)
	return nil
}

func (r repository) DeleteById(id valueobjects.Id) error {
	fmt.Println("CategoryRepositoryImpl.DeleteById")
	fmt.Println(id)
	return nil
}

func (r repository) GetAll() ([]entity.Category, error) {
	fmt.Println("CategoryRepositoryImpl.GetAll")
	return []entity.Category{}, nil
}

func (r repository) GetById(id valueobjects.Id) (entity.Category, error) {
	fmt.Println("CategoryRepositoryImpl.GetById")
	fmt.Println(id)
	return entity.Category{}, nil
}

func (r repository) Update(category entity.Category) error {
	fmt.Println("CategoryRepositoryImpl.Update")
	fmt.Println(category)
	return nil
}

func (r repository) GetSubcategoriesByCategoryId(id valueobjects.Id) ([]entity.Subcategory, error) {
	fmt.Println("CategoryRepositoryImpl.GetSubcategoriesByCategoryId")
	fmt.Println(id)
	return []entity.Subcategory{}, nil
}

func NewCategoryRepository() CategoryRepository {
	return repository{}
}
