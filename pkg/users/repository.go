package users

import (
	"log"

	"gorm.io/gorm"

	"gitlab.com/bpradana/privy-pretest/pkg/domain"
)

// Declare user repository struct
type UserRepository struct {
	db *gorm.DB
}

// Constructor function to create new user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	db.AutoMigrate(&domain.User{})

	return &UserRepository{
		db: db,
	}
}

// Function to create user in database
func (r *UserRepository) Create(user *domain.User) (*domain.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		log.Println("[users] [repository] error creating user, err: ", err.Error())
		return nil, err
	}

	return user, nil
}

// Function to get user by email
func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Println("[users] [repository] error getting user by email, err: ", err.Error())
		return nil, err
	}

	return &user, nil
}

// Function to get user by username
func (r *UserRepository) GetByUsername(username string) (*domain.User, error) {
	var user domain.User

	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Println("[users] [repository] error getting user by username, err: ", err.Error())
		return nil, err
	}

	return &user, nil
}
