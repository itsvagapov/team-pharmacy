package repository

import (
	"github.com/itsvagapov/team-pharmacy/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error

	GetAll() ([]models.Category, error)
}

type gormCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &gormCategoryRepository{db: db}
}

func (r *gormCategoryRepository) Create(category *models.Category) error {
	if category == nil {
		return nil
	}

	return r.db.Create(category).Error
}

func (r *gormCategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category

	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}
