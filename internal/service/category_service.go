package service

import (
	"errors"
	"strings"

	"github.com/itsvagapov/team-pharmacy/internal/models"
	"github.com/itsvagapov/team-pharmacy/internal/repository"
)

var ErrCategoryNameRequired = errors.New("имя не может быть пустым")

type CategoryService interface {
	CreateCategory(req models.CategoryCreateRequest) (*models.Category, error)

	GetAllCategories() ([]models.Category, error)
}

type categoryService struct {
	categories repository.CategoryRepository
}

func NewCategoryService(categories repository.CategoryRepository) CategoryService {
	return &categoryService{
		categories: categories,
	}
}

func (s *categoryService) CreateCategory(req models.CategoryCreateRequest) (*models.Category, error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, ErrCategoryNameRequired
	}

	category := &models.Category{
		Name: req.Name,
	}

	if err := s.categories.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	categories, err := s.categories.GetAll()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
