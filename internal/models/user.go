package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `json:"full_name" `
	Email    string `json:"email" `
	Phone    string `json:"phone" `
	DefaultAddress string `json:"default_adress"`
	
}

type UserCreateRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	DefaultAdress string `json:"default_adress" binding:"required"`
}





