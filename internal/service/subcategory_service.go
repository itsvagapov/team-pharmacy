package service

import (
	"errors"
	"strings"

	"github.com/itsvagapov/team-pharmacy/internal/models"
	"github.com/itsvagapov/team-pharmacy/internal/repository"
)

var ErrSubcategoryNameRequired = errors.New("имя не может быть пустым")
var ErrCategoryNotFound = errors.New("категория не найдена")

type SubcategoryService interface {
	CreateSubcategory(categoryID uint, req models.SubcategoryCreateRequest) (*models.Subcategory, error)
	GetSubcategoriesByCategoryID(categoryID uint) ([]models.Subcategory, error)
}

type subcategoryService struct {
	subcategories repository.SubcategoryRepository
	categories repository.CategoryRepository
}

func NewSubcategoryService(subcategories repository.SubcategoryRepository, categories repository.CategoryRepository) SubcategoryService {
	return &subcategoryService {
		categories: categories,
		subcategories: subcategories,
	}
}

func (s *subcategoryService) CreateSubcategory(categoryID uint, req models.SubcategoryCreateRequest) (*models.Subcategory, error) {

	if strings.TrimSpace(req.Name) == "" {
		return nil, ErrSubcategoryNameRequired
	}

	category, err := s.categories.GetByID(categoryID)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, ErrCategoryNotFound
	}

	subcategory := &models.Subcategory{
		CategoryID: categoryID,
		Name:       strings.TrimSpace(req.Name),
	}

	if err := s.subcategories.Create(subcategory); err != nil {
		return nil, err
	}

	return subcategory, nil
}

func (s *subcategoryService) GetSubcategoriesByCategoryID(categoryID uint) ([]models.Subcategory, error) {
	category, err := s.categories.GetByID(categoryID)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, ErrCategoryNotFound
	}

	subcategories, err := s.subcategories.GetByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}

	return subcategories, nil
}