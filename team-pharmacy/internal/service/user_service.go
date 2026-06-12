package service

import (
	"errors"

	"github.com/itsvagapov/team-pharmacy/internal/models"
	"github.com/itsvagapov/team-pharmacy/internal/repository"
)

type UserService interface {
	Create(userReq *models.UserCreate) error
	GetById(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) Create(userReq *models.UserCreate) error {

	//проверка email
	existingUser, err := s.userRepo.GetByEmail(userReq.Email)
	if existingUser != nil{
		return nil,errors.New("email already exists")
	}

	user := &models.User{
		FullName:      userReq.FullName,
		Email:         userReq.Email,
		Phone:         userReq.Phone,
		DefaultAdress: userReq.DefaultAdress,
	}

	err := s.userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}


