package users

import (
	"log"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"gitlab.com/bpradana/privy-pretest/cmd/auth"
	"gitlab.com/bpradana/privy-pretest/pkg/domain"
)

// Declare user usecase struct
type UserUsecase struct {
	userRepository    domain.UserRepository
	balanceRepository domain.UserBalanceRepository
}

// Constructor function to create new user usecase
func NewUserUsecase(userRepository domain.UserRepository, balanceRepository domain.UserBalanceRepository) *UserUsecase {
	return &UserUsecase{
		userRepository:    userRepository,
		balanceRepository: balanceRepository,
	}
}

// Function to register user
func (u *UserUsecase) Register(user *domain.User) (*domain.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[users] [usecase] [Register] error hashing password, err: ", err.Error())
		return nil, err
	}
	user.Password = string(hashedPassword)

	// Create user
	registeredUser, err := u.userRepository.Create(user)
	if err != nil {
		log.Println("[users] [usecase] [Register] error creating user, err: ", err.Error())
		return nil, err
	}

	// Create balance
	balanceToCreate := &domain.UserBalance{
		UserID:         registeredUser.ID,
		Balance:        0,
		BalanceAchieve: "Rp 0",
	}
	_, err = u.balanceRepository.Create(balanceToCreate)
	if err != nil {
		log.Println("[users] [usecase] [Register] error creating balance, err: ", err.Error())
		return nil, err
	}

	return registeredUser, nil
}

// Function to login user
func (u *UserUsecase) Login(user *domain.User, c echo.Context) (*domain.User, string, error) {
	// Get user by email
	userByEmail, err := u.userRepository.GetByEmail(user.Email)
	if err != nil {
		log.Println("[users] [usecase] [Login] error getting user by email, err: ", err.Error())
		return nil, "", err
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(userByEmail.Password), []byte(user.Password))
	if err != nil {
		log.Println("[users] [usecase] [Login] error comparing password, err: ", err.Error())
		return nil, "", err
	}

	// Generate token
	token, err := auth.GenerateTokenAndSetCookie(userByEmail, c)

	return userByEmail, token, nil
}

// Function to logout user
func (u *UserUsecase) Logout(c echo.Context) error {
	// Delete token cookie
	auth.DeleteTokenCookie(c)

	return nil
}
