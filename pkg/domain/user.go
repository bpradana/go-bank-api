package domain

import (
	"gorm.io/gorm"
)

// User is a struct that represent user table
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

// UserUseCase is a interface that represent user usecase
type UserUsecase interface {
	Register(u *User) (*User, error)
}

// UserRepository is a interface that represent user repository
type UserRepository interface {
	Create(u *User) (*User, error)
}
