package domain

import (
	"gorm.io/gorm"
)

// Balance is a struct that represent balance table
type Balance struct {
	gorm.Model
	UserID         uint   `json:"user_id" gorm:"unique"`
	Balance        int    `json:"balance"`
	BalanceAchieve string `json:"balance_achieve"`
	User           User   `json:"-"`
}

// BalanceUseCase is a interface that represent balance usecase
type BalanceUsecase interface {
	CheckBalance(username string) (*Balance, error)
}

// BalanceRepository is a interface that represent balance repository
type BalanceRepository interface {
	Create(b *Balance) (*Balance, error)
	GetByUserID(id uint) (*Balance, error)
}
