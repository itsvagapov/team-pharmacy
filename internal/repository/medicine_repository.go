package repository

import (
	"errors"

	"github.com/itsvagapov/team-pharmacy/internal/models"
	"gorm.io/gorm"
)

type MedicineRepository interface {
	Create(medicine *models.Medicine) error
	GetAll(filter models.MedicineFilter) ([]models.Medicine, error)
	GetByID(id uint) (*models.Medicine, error)
	Update(medicine *models.Medicine) error
	Delete(id uint) error
	UpdateAvgRating(id uint, avgRating float64) error
}

type gormMedicineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) MedicineRepository {
	return &gormMedicineRepository{db: db}
}

func (r *gormMedicineRepository) Create(medicine *models.Medicine) error {
	if medicine == nil {
		return errors.New("medicine is nill")
	}

	return r.db.Create(medicine).Error
}

func (r *gormMedicineRepository) GetAll(filter models.MedicineFilter) ([]models.Medicine, error) {
	var medicines []models.Medicine

	query := r.db.Model(&models.Medicine{})

	if filter.Search != "" {
		query = query.Where(
			"name ILIKE ? OR description ILIKE ?",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}

	if filter.CategoryID != nil {
		query = query.Where("category_id = ?", *filter.CategoryID)
	}

	if filter.SubcategoryID != nil {
		query = query.Where("subcategory_id = ?", *filter.SubcategoryID)
	}

	if filter.InStock != nil {
		query = query.Where("in_stock = ?", *filter.InStock)
	}

	err := query.Find(&medicines).Error
	if err != nil {
		return nil, err
	}

	return medicines, nil
}

func (r *gormMedicineRepository) GetByID(id uint) (*models.Medicine, error) {
	var medicine models.Medicine

	if err := r.db.First(&medicine, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &medicine, nil
}

func (r *gormMedicineRepository) Update(medicine *models.Medicine) error {
	if medicine == nil {
		return errors.New("medicine is nil")
	}

	return r.db.Save(medicine).Error
}

func (r *gormMedicineRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Medicine{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *gormMedicineRepository) UpdateAvgRating(id uint, avgRating float64) error {
	return r.db.Model(&models.Medicine{}).Where("id = ?", id).Update("avg_rating", avgRating).Error
}


