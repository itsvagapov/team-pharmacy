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
	reviews repository.ReviewRepository
}

func NewReviewService(reviews repository.ReviewRepository) ReviewService {
	return &reviewService{
		reviews: reviews,
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

	review := &models.Review{
		UserID:     req.UserID,
		MedicineID: medicineID,
		Rating:     req.Rating,
		Text:       reqTextTrimmed,
	}

	if err := s.reviews.Create(review); err != nil {
		return nil, err
	}

	return review, nil
}

func (s *reviewService) GetReviewsByMedicineID(medicineID uint) ([]models.Review, error) {
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

	if req.Text != nil {
		if strings.TrimSpace(*req.Text) == "" {
			return nil, ErrReviewTextRequired
		}

		review.Text = strings.TrimSpace(*req.Text)
	}

	if err := s.reviews.Update(review); err != nil {
		return nil, err
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

	if err := s.reviews.Delete(id); err != nil {
		return err
	}

	return nil
}
