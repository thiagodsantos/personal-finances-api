package usecase

import (
	"fmt"

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

	exists, err := cc.repo.GetByName(category.Name)
	if err != nil {
		return err
	}

	if exists.Id != "" {
		return fmt.Errorf("category already exists with name %s", category.Name)
	}

	err = cc.repo.Create(category)
	if err != nil {
		return err
	}

	return nil
}
