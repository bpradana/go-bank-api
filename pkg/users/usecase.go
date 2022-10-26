package users

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"gitlab.com/bpradana/privy-pretest/cmd/auth"
	"gitlab.com/bpradana/privy-pretest/pkg/domain"
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
	registeredUser, err := u.userRepository.Create(user)
	if err != nil {
		log.Println("[users] [usecase] error creating user, err: ", err.Error())
		return nil, err
	}

	return registeredUser, nil
}

// Function to login user
func (u *UserUsecase) Login(user *domain.User) (*domain.User, string, error) {
	// Get user by email
	userByEmail, err := u.userRepository.GetByEmail(user.Email)
	if err != nil {
		log.Println("[users] [usecase] error getting user by email, err: ", err.Error())
		return nil, "", err
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(userByEmail.Password), []byte(user.Password))
	if err != nil {
		log.Println("[users] [usecase] error comparing password, err: ", err.Error())
		return nil, "", err
	}

	// Generate token
	token, err := auth.GenerateToken(userByEmail)

	return userByEmail, token, nil
}
