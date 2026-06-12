package repository

import (
	"errors"

	"github.com/itsvagapov/team-pharmacy/internal/models"
	"gorm.io/gorm"
)

type SubcategoryRepository interface {
	Create(subcategory *models.Subcategory) error
	GetByCategoryID(categoryID uint) ([]models.Subcategory, error)
}

type gormSubcategoryRepository struct {
	db *gorm.DB
}

func NewSubcategoryRepository(db *gorm.DB) SubcategoryRepository {
	return &gormSubcategoryRepository{db: db}
}

func (r *gormSubcategoryRepository) Create(subcategory *models.Subcategory) error {
	if subcategory == nil {
		return errors.New("subcategory is nill")
	}

	return r.db.Create(subcategory).Error
}

func (r *gormSubcategoryRepository) GetByCategoryID(categoryID uint) ([]models.Subcategory, error) {
	var subcategories []models.Subcategory

	err := r.db.Where("category_id = ?", categoryID).Find(&subcategories).Error
	if err != nil {
		return nil, err
	}

	return subcategories, nil
}
