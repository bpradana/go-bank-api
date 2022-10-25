package users

import (
	"log"

	"gitlab.com/bpradana/privy-pretest/pkg/domain"
	"golang.org/x/crypto/bcrypt"
)

// Declare user usecase struct
type UserUsecase struct {
	userRepository domain.UserRepository
}

// Constructor function to create new user usecase
func NewUserUsecase(userRepository domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

// Function to register user
func (u *UserUsecase) Register(user *domain.User) (*domain.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[users] [usecase] error hashing password, err: ", err.Error())
		return nil, err
	}
	user.Password = string(hashedPassword)

	// Create user
	_, err = u.userRepository.Create(user)
	if err != nil {
		log.Println("[users] [usecase] error creating user, err: ", err.Error())
		return nil, err
	}

	return user, nil
}
