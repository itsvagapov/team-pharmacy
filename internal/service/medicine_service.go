package service

import (
	"errors"
	"strings"

	"github.com/itsvagapov/team-pharmacy/internal/models"
	"github.com/itsvagapov/team-pharmacy/internal/repository"
)

var ErrMedicineNotFound = errors.New("лекарство не найдено")
var ErrMedicineNameRequired = errors.New("имя не может быть пустым")
var ErrMedicineDescriptionRequired = errors.New("описание не может быть пустым")
var ErrManufacturerRequired = errors.New("производитель не может быть пустым")
var ErrPriceMustBePositive = errors.New("цена должна быть больше нуля")
var ErrStockQuantityNegative = errors.New("количество не может быть отрицательным")
var ErrSubcategoryNotFound = errors.New("подкатегория не найдена")

type MedicineService interface {
	CreateMedicine(req models.MedicineCreateRequest) (*models.Medicine, error)
	GetAllMedicines(filter models.MedicineFilter) ([]models.Medicine, error)
	GetMedicineByID(id uint) (*models.Medicine, error)
	UpdateMedicine(id uint, req models.MedicineUpdateRequest) (*models.Medicine, error)
	DeleteMedicine(id uint) error
}

type medicineService struct {
	medicines     repository.MedicineRepository
	categories    repository.CategoryRepository
	subcategories repository.SubcategoryRepository
}

func NewMedicineService(medicines repository.MedicineRepository, categories repository.CategoryRepository, subcategories repository.SubcategoryRepository) MedicineService {
	return &medicineService{
		medicines:     medicines,
		categories:    categories,
		subcategories: subcategories,
	}
}

func (s *medicineService) CreateMedicine(req models.MedicineCreateRequest) (*models.Medicine, error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, ErrMedicineNameRequired
	}

	if strings.TrimSpace(req.Description) == "" {
		return nil, ErrMedicineDescriptionRequired
	}

	if strings.TrimSpace(req.Manufacturer) == "" {
		return nil, ErrManufacturerRequired
	}

	if req.Price <= 0 {
		return nil, ErrPriceMustBePositive
	}

	if req.StockQuantity < 0 {
		return nil, ErrStockQuantityNegative
	}

	category, err := s.categories.GetByID(req.CategoryID)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, ErrCategoryNotFound
	}

	subcategories, err := s.subcategories.GetByCategoryID(req.CategoryID)
	if err != nil {
		return nil, err
	}

	subcategoryExists := false

	for _, subcategory := range subcategories {
		if subcategory.ID == req.SubcategoryID {
			subcategoryExists = true
			break
		}
	}

	if !subcategoryExists {
		return nil, ErrSubcategoryNotFound
	}

	medicine := &models.Medicine{
		Name:                 strings.TrimSpace(req.Name),
		Description:          strings.TrimSpace(req.Description),
		Price:                req.Price,
		InStock:              req.InStock,
		StockQuantity:        req.StockQuantity,
		CategoryID:           req.CategoryID,
		SubcategoryID:        req.SubcategoryID,
		Manufacturer:         strings.TrimSpace(req.Manufacturer),
		PrescriptionRequired: req.PrescriptionRequired,
		AvgRating:            0,
	}

	if err := s.medicines.Create(medicine); err != nil {
		return nil, err
	}

	return medicine, nil
}

func (s *medicineService) GetAllMedicines(filter models.MedicineFilter) ([]models.Medicine, error) {
	medicines, err := s.medicines.GetAll(filter)
	if err != nil {
		return nil, err
	}

	return medicines, nil
}

func (s *medicineService) GetMedicineByID(id uint) (*models.Medicine, error) {
	medicine, err := s.medicines.GetByID(id)
	if err != nil {
		return nil, err
	}

	if medicine == nil {
		return nil, ErrMedicineNotFound
	}

	return medicine, nil
}

func (s *medicineService) UpdateMedicine(id uint, req models.MedicineUpdateRequest) (*models.Medicine, error) {
	medicine, err := s.medicines.GetByID(id)
	if err != nil {
		return nil, err
	}

	if medicine == nil {
		return nil, ErrMedicineNotFound
	}

	if req.Price != nil {
		if *req.Price <= 0 {
			return nil, ErrPriceMustBePositive
		}

		medicine.Price = *req.Price
	}

	if req.StockQuantity != nil {
		if *req.StockQuantity < 0 {
			return nil, ErrStockQuantityNegative
		}

		medicine.StockQuantity = *req.StockQuantity
	}

	if req.InStock != nil {
		medicine.InStock = *req.InStock
	}

	if err := s.medicines.Update(medicine); err != nil {
		return nil, err
	}

	return medicine, nil
}

func (s *medicineService) DeleteMedicine(id uint) error {
	medicine, err := s.medicines.GetByID(id)
	if err != nil {
		return err
	}

	if medicine == nil {
		return ErrMedicineNotFound
	}

	if err := s.medicines.Delete(id); err != nil {
		return err
	}

	return nil
}
