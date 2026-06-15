package service

import (
	"errors"
	"fmt"

	"github.com/itsvagapov/team-pharmacy/internal/models"
	"github.com/itsvagapov/team-pharmacy/internal/repository"
	"gorm.io/gorm"
)

type UserService interface {
	Create(userReq *models.UserCreateRequest) (*models.User, error)
	GetById(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Create(userReq *models.UserCreateRequest) (*models.User, error) {
	//проверка email
	existingUser, err := s.userRepo.GetByEmail(userReq.Email)
	if err != nil {
		
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("database check error: %w", err)
		}
	} else if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	newUser := &models.User{
		FullName:       userReq.FullName,
		Email:          userReq.Email,
		Phone:          userReq.Phone,
		DefaultAddress: userReq.DefaultAdress,
	}

	err = s.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *userService) GetById(id uint) (*models.User, error) {

	return s.userRepo.GetByID(id)
}

func (s *userService) GetByEmail(email string) (*models.User, error) {
	// Вызываем метод репозитория и возвращаем результат как есть
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
