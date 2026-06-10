package service

import (
	"errors"

	"github.com/itsvagapov/team-pharmacy/internal/models"
	"github.com/itsvagapov/team-pharmacy/internal/repository"
)

var ErrCategoryNotFound = errors.New("категория не найдена")

type CategoryService interface {
	CreateCategory(req models.CreateCategoryRequest) (*models.Category, error)

	GetAllCategories() ([]models.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) CreateCategory(req models.CreateCategoryRequest) (*models.Category, error) {
	category := &models.Category{
		Name: req.Name,
	}

	err := s.repo.Create(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	categories, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
