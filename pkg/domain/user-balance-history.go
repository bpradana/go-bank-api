package domain

import (
	"gorm.io/gorm"
)

type Type string

const (
	Kredit Type = "kredit"
	Debit  Type = "debit"
)

// UserBalanceHistory is a struct that represent user balance history table
type UserBalanceHistory struct {
	gorm.Model
	UserBalanceID uint   `json:"user_balance_id"`
	BalanceBefore int    `json:"balance_before"`
	BalanceAfter  int    `json:"balance_after"`
	Activity      string `json:"activity"`
	Type          Type   `json:"type" sql:"type:ENUM('kredit', 'debit')"`
	Ip            string `json:"ip"`
	Location      string `json:"location"`
	UserAgent     string `json:"user_agent"`
	Author        string `json:"author"`
}

// UserBalanceHistoryUseCase is a interface that represent user balance history usecase
type UserBalanceHistoryUsecase interface {
	Check(username string) (*[]UserBalanceHistory, error)
}

// UserBalanceHistoryRepository is a interface that represent user balance history repository
type UserBalanceHistoryRepository interface {
	GetByBalanceID(id uint) (*[]UserBalanceHistory, error)
	Create(userBalanceHistory *UserBalanceHistory) (*UserBalanceHistory, error)
}
