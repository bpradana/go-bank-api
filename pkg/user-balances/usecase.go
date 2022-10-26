package userBalances

import (
	"errors"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"gitlab.com/bpradana/privy-pretest/cmd/util"
	"gitlab.com/bpradana/privy-pretest/pkg/domain"
)

// Declare balance usecase struct
type BalanceUsecase struct {
	userRepository               domain.UserRepository
	userBalanceRepository        domain.UserBalanceRepository
	userBalanceHistoryRepository domain.UserBalanceHistoryRepository
}

// Constructor function to create new balance usecase
func NewUserBalanceUsecase(
	userRepository domain.UserRepository,
	userBalanceRepository domain.UserBalanceRepository,
	userBalanceHistoryRepository domain.UserBalanceHistoryRepository,
) *BalanceUsecase {
	return &BalanceUsecase{
		userRepository:               userRepository,
		userBalanceRepository:        userBalanceRepository,
		userBalanceHistoryRepository: userBalanceHistoryRepository,
	}
}

// Function to get balance by username
func (u *BalanceUsecase) Check(username string) (*domain.UserBalance, error) {
	// Get user by username
	user, err := u.userRepository.GetByUsername(username)
	if err != nil {
		log.Println("[balances] [usecase] error getting user by username, err: ", err.Error())
		return nil, err
	}

	// Get balance by user id
	balance, err := u.userBalanceRepository.GetByUserID(user.ID)
	if err != nil {
		log.Println("[balances] [usecase] error getting balance by username, err: ", err.Error())
		return nil, err
	}

	return balance, nil
}

// Function to deposit balance
func (u *BalanceUsecase) Deposit(username string, amount int, c echo.Context) (*domain.UserBalance, error) {
	// Get user by username
	user, err := u.userRepository.GetByUsername(username)
	if err != nil {
		log.Println("[balances] [usecase] error getting user by username, err: ", err.Error())
		return nil, err
	}

	// Get balance by user id
	balance, err := u.userBalanceRepository.GetByUserID(user.ID)
	if err != nil {
		log.Println("[balances] [usecase] error getting balance by username, err: ", err.Error())
		return nil, err
	}

	// Update balance
	balanceBefore := balance.Balance
	balance.Balance += amount
	balance.BalanceAchieve = fmt.Sprintf("Rp. %d", balance.Balance)
	balance, err = u.userBalanceRepository.Update(balance.ID, balance)
	if err != nil {
		log.Println("[balances] [usecase] error updating balance, err: ", err.Error())
		return nil, err
	}

	// Add transaction history
	ip := c.RealIP()
	userAgent := c.Request().UserAgent()
	location, err := util.GetIPLocation(ip)
	if err != nil {
		log.Println("[balances] [usecase] error getting ip location, err: ", err.Error())
		return nil, err
	}

	transaction := &domain.UserBalanceHistory{
		UserBalanceID: balance.ID,
		BalanceBefore: balanceBefore,
		BalanceAfter:  balance.Balance,
		Activity:      "deposit",
		Type:          domain.Kredit,
		Ip:            ip,
		Location:      location,
		UserAgent:     userAgent,
		Author:        user.Username,
	}
	_, err = u.userBalanceHistoryRepository.Create(transaction)
	if err != nil {
		log.Println("[balances] [usecase] error creating transaction history, err: ", err.Error())
		return nil, err
	}

	return balance, nil
}

// Function to withdraw balance
func (u *BalanceUsecase) Withdraw(username string, amount int, c echo.Context) (*domain.UserBalance, error) {
	// Get user by username
	user, err := u.userRepository.GetByUsername(username)
	if err != nil {
		log.Println("[balances] [usecase] error getting user by username, err: ", err.Error())
		return nil, err
	}

	// Get balance by user id
	balance, err := u.userBalanceRepository.GetByUserID(user.ID)
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
	balanceBefore := balance.Balance
	balance.Balance -= amount
	balance.BalanceAchieve = fmt.Sprintf("Rp. %d", balance.Balance)
	balance, err = u.userBalanceRepository.Update(balance.ID, balance)
	if err != nil {
		log.Println("[balances] [usecase] error updating balance, err: ", err.Error())
		return nil, err
	}

	// Add transaction history
	ip := c.RealIP()
	userAgent := c.Request().UserAgent()
	location, err := util.GetIPLocation(ip)
	if err != nil {
		log.Println("[balances] [usecase] error getting ip location, err: ", err.Error())
		return nil, err
	}

	transaction := &domain.UserBalanceHistory{
		UserBalanceID: balance.ID,
		BalanceBefore: balanceBefore,
		BalanceAfter:  balance.Balance,
		Activity:      "withdraw",
		Type:          domain.Debit,
		Ip:            ip,
		Location:      location,
		UserAgent:     userAgent,
		Author:        user.Username,
	}
	_, err = u.userBalanceHistoryRepository.Create(transaction)
	if err != nil {
		log.Println("[balances] [usecase] error creating transaction history, err: ", err.Error())
		return nil, err
	}

	return balance, nil
}

// Function to transfer balance
func (u *BalanceUsecase) Transfer(username string, amount int, toUsername string, c echo.Context) (*domain.UserBalance, error) {
	// Get user by username
	user, err := u.userRepository.GetByUsername(username)
	if err != nil {
		log.Println("[balances] [usecase] error getting user by username, err: ", err.Error())
		return nil, err
	}

	// Get balance by user id
	balance, err := u.userBalanceRepository.GetByUserID(user.ID)
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
	toBalance, err := u.userBalanceRepository.GetByUserID(toUser.ID)
	if err != nil {
		log.Println("[balances] [usecase] error getting balance by username, err: ", err.Error())
		return nil, err
	}

	// Update balance
	balanceBefore := balance.Balance
	balance.Balance -= amount
	balance.BalanceAchieve = fmt.Sprintf("Rp. %d", balance.Balance)
	balance, err = u.userBalanceRepository.Update(balance.ID, balance)
	if err != nil {
		log.Println("[balances] [usecase] error updating balance, err: ", err.Error())
		return nil, err
	}

	// Add transaction history
	ip := c.RealIP()
	userAgent := c.Request().UserAgent()
	location, err := util.GetIPLocation(ip)
	if err != nil {
		log.Println("[balances] [usecase] error getting ip location, err: ", err.Error())
		return nil, err
	}

	transaction := &domain.UserBalanceHistory{
		UserBalanceID: balance.ID,
		BalanceBefore: balanceBefore,
		BalanceAfter:  balance.Balance,
		Activity:      "transfer",
		Type:          domain.Kredit,
		Ip:            ip,
		Location:      location,
		UserAgent:     userAgent,
		Author:        user.Username,
	}
	_, err = u.userBalanceHistoryRepository.Create(transaction)
	if err != nil {
		log.Println("[balances] [usecase] error creating transaction history, err: ", err.Error())
		return nil, err
	}

	// Update balance
	toBalanceBefore := toBalance.Balance
	toBalance.Balance += amount
	toBalance.BalanceAchieve = fmt.Sprintf("Rp. %d", toBalance.Balance)
	toBalance, err = u.userBalanceRepository.Update(toBalance.ID, toBalance)
	if err != nil {
		log.Println("[balances] [usecase] error updating balance, err: ", err.Error())
		return nil, err
	}

	// Add transaction history
	toTransaction := &domain.UserBalanceHistory{
		UserBalanceID: toBalance.ID,
		BalanceBefore: toBalanceBefore,
		BalanceAfter:  toBalance.Balance,
		Activity:      "transfer",
		Type:          domain.Debit,
		Ip:            ip,
		Location:      location,
		UserAgent:     userAgent,
		Author:        user.Username,
	}
	_, err = u.userBalanceHistoryRepository.Create(toTransaction)
	if err != nil {
		log.Println("[balances] [usecase] error creating transaction history, err: ", err.Error())
		return nil, err
	}

	return balance, nil
}
