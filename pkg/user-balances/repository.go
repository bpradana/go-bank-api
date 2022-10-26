package userBalances

import (
	"log"

	"gorm.io/gorm"

	"gitlab.com/bpradana/privy-pretest/pkg/domain"
)

// Declare balance repository struct
type BalanceRepository struct {
	db *gorm.DB
}

// Constructor function to create new balance repository
func NewUserBalanceRepository(db *gorm.DB) *BalanceRepository {
	db.AutoMigrate(&domain.UserBalance{})

	return &BalanceRepository{
		db: db,
	}
}

// Function to create balance in database
func (r *BalanceRepository) Create(balance *domain.UserBalance) (*domain.UserBalance, error) {
	err := r.db.Create(balance).Error
	if err != nil {
		log.Println("[balances] [repository] error creating balance, err: ", err.Error())
		return nil, err
	}

	return balance, nil
}

// Function to get balance by user id
func (r *BalanceRepository) GetByUserID(id uint) (*domain.UserBalance, error) {
	var balance domain.UserBalance

	err := r.db.Where("user_id = ?", id).First(&balance).Error
	if err != nil {
		log.Println("[balances] [repository] [GetByUserID] error getting balance by user id, err: ", err.Error())
		return nil, err
	}

	return &balance, nil
}

// Function to update balance in database
func (r *BalanceRepository) Update(id uint, balance *domain.UserBalance) (*domain.UserBalance, error) {
	var oldBalance domain.UserBalance

	err := r.db.Where("id = ?", id).First(&oldBalance).Error
	if err != nil {
		log.Println("[balances] [repository] [Update] error getting balance by id, err: ", err.Error())
		return nil, err
	}

	err = r.db.Model(&oldBalance).Updates(balance).Error
	if err != nil {
		log.Println("[balances] [repository] [Update] error updating balance, err: ", err.Error())
		return nil, err
	}

	return balance, nil
}
