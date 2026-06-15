package repository

import (
	"errors"

	"github.com/itsvagapov/team-pharmacy/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	if user == nil{
		return errors.New("user is nil")
	}
	err := r.db.Create(user).Error
	if err != nil{
		return err
	}
	
	return nil
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	
	// Ищем первую запись, где email совпадает
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		// Возвращаем nil и саму ошибку (GORM вернет gorm.ErrRecordNotFound, если юзера нет)
		return nil, err 
	}
	
	return &user, nil
}


