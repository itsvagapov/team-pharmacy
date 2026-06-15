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
}

type userService struct {
	userRepo repository.UserRepository
}	

func NewUserService(userRepo repository.UserRepository) UserService{
	return &userService{userRepo: userRepo}
}

func (s *userService) Create(userReq *models.UserCreate) error {

	//проверка email
	existingUser, _ := s.userRepo.GetByEmail(userReq.Email)
	if existingUser != nil{
		return errors.New("email already exists")
	}

	user := &models.User{
		FullName:      userReq.FullName,
		Email:         userReq.Email,
		Phone:         userReq.Phone,
		DefaultAdress: userReq.DefaultAdress,
	}

	_, err := s.userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetById(id uint) (*models.User,error) {

	return s.userRepo.GetByID(id)
}

func (s *userService) GetByEmail(email string) (*models.User, error){
	return nil,nil
}


