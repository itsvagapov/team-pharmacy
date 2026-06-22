package service

import (
	"errors"
	"strings"
	"github.com/itsvagapov/team-pharmacy/internal/models"
	"github.com/itsvagapov/team-pharmacy/internal/repository"
)

var ErrReviewNotFound = errors.New("отзыв не найден")
var ErrRatingOutOfRange = errors.New("оценка должна быть от 1 до 5")
var ErrReviewTextRequired = errors.New("текст отзыва не может быть пустым")

type ReviewService interface {
	CreateReview(medicineID uint, req models.ReviewCreateRequest) (*models.Review, error)
	GetReviewsByMedicineID(medicineID uint) ([]models.Review, error)
	UpdateReview(id uint, req models.ReviewUpdateRequest) (*models.Review, error)
	DeleteReview(id uint) error
}

type reviewService struct {
	reviews   repository.ReviewRepository
	medicines repository.MedicineRepository
}

func NewReviewService(reviews repository.ReviewRepository, medicines repository.MedicineRepository) ReviewService {
	return &reviewService{
		reviews:   reviews,
		medicines: medicines,
	}
}

func (s *reviewService) CreateReview(medicineID uint, req models.ReviewCreateRequest) (*models.Review, error) {
	if req.Rating < 1 || req.Rating > 5 {
		return nil, ErrRatingOutOfRange
	}

	reqTextTrimmed := strings.TrimSpace(req.Text)

	if reqTextTrimmed == "" {
		return nil, ErrReviewTextRequired
	}

	medicine, err := s.medicines.GetByID(medicineID)
	if err != nil {
		return nil, err
	}

	if medicine == nil {
		return nil, ErrMedicineNotFound
	}

	review := &models.Review{
		UserID:     req.UserID,
		MedicineID: medicineID,
		Rating:     req.Rating,
		Text:       reqTextTrimmed,
	}

	if err := s.reviews.Create(review); err != nil {
		return nil, err
	}

	avgRating, err := s.reviews.GetAverageRatingByMedicineID(medicineID)
	if err != nil {
		return nil, err
	}

	if err := s.medicines.UpdateAvgRating(medicineID, avgRating); err != nil {
		return nil, err
	}

	return review, nil
}

func (s *reviewService) GetReviewsByMedicineID(medicineID uint) ([]models.Review, error) {
	medicine, err := s.medicines.GetByID(medicineID)
	if err != nil {
		return nil, err
	}

	if medicine == nil {
		return nil, ErrMedicineNotFound
	}

	reviews, err := s.reviews.GetByMedicineID(medicineID)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (s *reviewService) UpdateReview(id uint, req models.ReviewUpdateRequest) (*models.Review, error) {
	review, err := s.reviews.GetByID(id)
	if err != nil {
		return nil, err
	}

	if review == nil {
		return nil, ErrReviewNotFound
	}

	ratingChanged := false

	if req.Rating != nil {
		if *req.Rating < 1 || *req.Rating > 5 {
			return nil, ErrRatingOutOfRange
		}

		review.Rating = *req.Rating
		ratingChanged = true
	}

	if req.Text != nil {
		if strings.TrimSpace(*req.Text) == "" {
			return nil, ErrReviewTextRequired
		}

		review.Text = strings.TrimSpace(*req.Text)
	}

	if err := s.reviews.Update(review); err != nil {
		return nil, err
	}

	if ratingChanged {
		avgRating, err := s.reviews.GetAverageRatingByMedicineID(review.MedicineID)
		if err != nil {
			return nil, err
		}

		if err := s.medicines.UpdateAvgRating(review.MedicineID, avgRating); err != nil {
			return nil, err
		}
	}

	return review, nil
}

func (s *reviewService) DeleteReview(id uint) error {
	review, err := s.reviews.GetByID(id)
	if err != nil {
		return err
	}

	if review == nil {
		return ErrReviewNotFound
	}

	medicineID := review.MedicineID

	if err := s.reviews.Delete(id); err != nil {
		return err
	}

	avgRating, err := s.reviews.GetAverageRatingByMedicineID(medicineID)
	if err != nil {
		return err
	}

	if err := s.medicines.UpdateAvgRating(medicineID, avgRating); err != nil {
		return err
	}

	return nil
}
