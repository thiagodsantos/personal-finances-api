package usecase

import (
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/category/entity"
	"github.com/thiagodsantos/personal-finances-api/pkg/domain/category/repository"
	"github.com/thiagodsantos/personal-finances-api/pkg/valueobjects"
)

type CreateCategoryInput struct {
	Description valueobjects.Text
	Name        valueobjects.Name
}

type CreateCategory struct {
	repo repository.CategoryRepository
}

func (cc *CreateCategory) Execute(input CreateCategoryInput) error {
	category := entity.Category{
		Description: input.Description,
		Name:        input.Name,
	}

	err := category.IsValid()
	if err != nil {
		return err
	}

	err = cc.repo.Create(category)
	if err != nil {
		return err
	}

	return nil
}
