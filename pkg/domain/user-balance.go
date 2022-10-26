package domain

import (
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

// UserBalance is a struct that represent user balance table
type UserBalance struct {
	gorm.Model
	UserID               uint                 `json:"user_id" gorm:"unique"`
	Balance              int                  `json:"balance"`
	BalanceAchieve       string               `json:"balance_achieve"`
	User                 User                 `json:"-"`
	UserBalanceHistories []UserBalanceHistory `json:"-"`
}

// UserBalanceUseCase is a interface that represent user balance usecase
type UserBalanceUsecase interface {
	Check(username string) (*UserBalance, error)
	Deposit(username string, amount int, c echo.Context) (*UserBalance, error)
	Withdraw(username string, amount int, c echo.Context) (*UserBalance, error)
	Transfer(username string, amount int, toUsername string, c echo.Context) (*UserBalance, error)
}

// UserBalanceRepository is a interface that represent user balance repository
type UserBalanceRepository interface {
	GetByUserID(id uint) (*UserBalance, error)
	Create(b *UserBalance) (*UserBalance, error)
	Update(id uint, b *UserBalance) (*UserBalance, error)
}
