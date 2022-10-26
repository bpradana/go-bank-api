package balances

import (
	"log"

	"gitlab.com/bpradana/privy-pretest/pkg/domain"
)

// Declare balance usecase struct
type BalanceUsecase struct {
	balanceRepository domain.BalanceRepository
	userRepository    domain.UserRepository
}

// Constructor function to create new balance usecase
func NewBalanceUsecase(balanceRepository domain.BalanceRepository, userRepository domain.UserRepository) *BalanceUsecase {
	return &BalanceUsecase{
		balanceRepository: balanceRepository,
		userRepository:    userRepository,
	}
}

// Function to get balance by username
func (u *BalanceUsecase) CheckBalance(username string) (*domain.Balance, error) {
	// Get user by username
	user, err := u.userRepository.GetByUsername(username)
	if err != nil {
		log.Println("[balances] [usecase] error getting user by username, err: ", err.Error())
		return nil, err
	}

	// Get balance by user id
	balance, err := u.balanceRepository.GetByUserID(user.ID)
	if err != nil {
		log.Println("[balances] [usecase] error getting balance by username, err: ", err.Error())
		return nil, err
	}

	return balance, nil
}
