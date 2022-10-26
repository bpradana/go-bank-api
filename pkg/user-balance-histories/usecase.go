package userBalanceHistories

import (
	"log"

	"gitlab.com/bpradana/privy-pretest/pkg/domain"
)

// Declare user balance history usecase struct
type UserBalanceHistoryUsecase struct {
	userRepository               domain.UserRepository
	userBalanceRepository        domain.UserBalanceRepository
	userBalanceHistoryRepository domain.UserBalanceHistoryRepository
}

// Constructor function to create new user balance history usecase
func NewUserBalanceHistoryUsecase(
	userRepository domain.UserRepository,
	userBalanceRepository domain.UserBalanceRepository,
	userBalanceHistoryRepository domain.UserBalanceHistoryRepository,
) *UserBalanceHistoryUsecase {
	return &UserBalanceHistoryUsecase{
		userRepository:               userRepository,
		userBalanceRepository:        userBalanceRepository,
		userBalanceHistoryRepository: userBalanceHistoryRepository,
	}
}

// Function to get user balance history by username
func (u *UserBalanceHistoryUsecase) Check(username string) (*[]domain.UserBalanceHistory, error) {
	// Get user by username
	user, err := u.userRepository.GetByUsername(username)
	if err != nil {
		log.Println("[balance history] [usecase] [Check] error getting user by username, err: ", err.Error())
		return nil, err
	}

	// Get balance by user id
	balance, err := u.userBalanceRepository.GetByUserID(user.ID)
	if err != nil {
		log.Println("[balance history] [usecase] [GetByUserID] error getting balance by username, err: ", err.Error())
		return nil, err
	}

	// Get balance history by balance id
	histories, err := u.userBalanceHistoryRepository.GetByBalanceID(balance.ID)
	if err != nil {
		log.Println("[balance history] [usecase] [GetByBalanceID] error getting balance history by balance id, err: ", err.Error())
		return nil, err
	}

	return histories, nil
}
