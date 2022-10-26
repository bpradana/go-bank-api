package userBalanceHistories

import (
	"log"

	"gorm.io/gorm"

	"gitlab.com/bpradana/privy-pretest/pkg/domain"
)

// Declare user balance history repository struct
type UserBalanceHistoryRepository struct {
	db *gorm.DB
}

// Constructor function to create new user balance history repository
func NewUserBalanceHistoryRepository(db *gorm.DB) *UserBalanceHistoryRepository {
	db.AutoMigrate(&domain.UserBalanceHistory{})

	return &UserBalanceHistoryRepository{
		db: db,
	}
}

// Function to get user balance history by balance id in database
func (r *UserBalanceHistoryRepository) GetByBalanceID(id uint) (*[]domain.UserBalanceHistory, error) {
	var histories []domain.UserBalanceHistory

	err := r.db.Where("user_balance_id = ?", id).Find(&histories).Error
	if err != nil {
		log.Println("[balance history] [repository] [GetByBalanceID] error getting balance history by balance id, err: ", err.Error())
		return nil, err
	}

	return &histories, nil
}

// Function to create new user balance history in database
func (r *UserBalanceHistoryRepository) Create(userBalanceHistory *domain.UserBalanceHistory) (*domain.UserBalanceHistory, error) {
	err := r.db.Create(userBalanceHistory).Error
	if err != nil {
		log.Println("[balance history] [repository] [Create] error creating new balance history, err: ", err.Error())
		return nil, err
	}

	return userBalanceHistory, nil
}
