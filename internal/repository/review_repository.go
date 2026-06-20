package repository

import (
	"errors"

	"github.com/itsvagapov/team-pharmacy/internal/models"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	Create(review *models.Review) error
	GetByID(id uint) (*models.Review, error)
	GetByMedicineID(medicineID uint) ([]models.Review, error)
	Update(review *models.Review) error
	Delete(id uint) error
	GetAverageRatingByMedicineID(medicineID uint) (float64, error)
}

type gormReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &gormReviewRepository{db: db}
}

func (r *gormReviewRepository) Create(review *models.Review) error {
	if review == nil {
		return errors.New("review is nil")
	}

	return r.db.Create(review).Error
}

func (r *gormReviewRepository) GetByID(id uint) (*models.Review, error) {
	var review models.Review

	if err := r.db.First(&review, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &review, nil
}

func (r *gormReviewRepository) GetByMedicineID(medicineID uint) ([]models.Review, error) {
	var reviews []models.Review

	err := r.db.
		Where("medicine_id = ?", medicineID).
		Find(&reviews).
		Error

	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *gormReviewRepository) Update(review *models.Review) error {
	if review == nil {
		return errors.New("review is nil")
	}

	return r.db.Save(review).Error
}

func (r *gormReviewRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Review{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *gormReviewRepository) GetAverageRatingByMedicineID(medicineID uint) (float64, error) {
	var avgRating float64

	err := r.db.
		Model(&models.Review{}).
		Where("medicine_id = ?", medicineID).
		Select("AVG(rating)").
		Scan(&avgRating).
		Error

	if err != nil {
		return 0, err
	}

	return avgRating, nil
}
