package balances

import (
	"errors"
	"fmt"
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

// Function to deposit balance
func (u *BalanceUsecase) Deposit(username string, amount int) (*domain.Balance, error) {
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

	// Update balance
	balance.Balance += amount
	balance.BalanceAchieve = fmt.Sprintf("Rp. %d", balance.Balance)
	balance, err = u.balanceRepository.Update(balance.ID, balance)
	if err != nil {
		log.Println("[balances] [usecase] error updating balance, err: ", err.Error())
		return nil, err
	}

	return balance, nil
}

// Function to withdraw balance
func (u *BalanceUsecase) Withdraw(username string, amount int) (*domain.Balance, error) {
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

	// Check if balance is enough
	if balance.Balance < amount {
		log.Println("[balances] [usecase] balance is not enough")
		return nil, errors.New("balance is not enough")
	}

	// Update balance
	balance.Balance -= amount
	balance.BalanceAchieve = fmt.Sprintf("Rp. %d", balance.Balance)
	balance, err = u.balanceRepository.Update(balance.ID, balance)
	if err != nil {
		log.Println("[balances] [usecase] error updating balance, err: ", err.Error())
		return nil, err
	}

	return balance, nil
}

// Function to transfer balance
func (u *BalanceUsecase) Transfer(username string, amount int, toUsername string) (*domain.Balance, error) {
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

	// Check if balance is enough
	if balance.Balance < amount {
		log.Println("[balances] [usecase] balance is not enough")
		return nil, errors.New("balance is not enough")
	}

	// Get user by username
	toUser, err := u.userRepository.GetByUsername(toUsername)
	if err != nil {
		log.Println("[balances] [usecase] error getting user by username, err: ", err.Error())
		return nil, err
	}

	// Get balance by user id
	toBalance, err := u.balanceRepository.GetByUserID(toUser.ID)
	if err != nil {
		log.Println("[balances] [usecase] error getting balance by username, err: ", err.Error())
		return nil, err
	}

	// Update balance
	balance.Balance -= amount
	balance.BalanceAchieve = fmt.Sprintf("Rp. %d", balance.Balance)
	balance, err = u.balanceRepository.Update(balance.ID, balance)
	if err != nil {
		log.Println("[balances] [usecase] error updating balance, err: ", err.Error())
		return nil, err
	}

	// Update balance
	toBalance.Balance += amount
	toBalance.BalanceAchieve = fmt.Sprintf("Rp. %d", toBalance.Balance)
	toBalance, err = u.balanceRepository.Update(toBalance.ID, toBalance)
	if err != nil {
		log.Println("[balances] [usecase] error updating balance, err: ", err.Error())
		return nil, err
	}

	return balance, nil
}
