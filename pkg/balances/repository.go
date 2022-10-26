package balances

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
func NewBalanceRepository(db *gorm.DB) *BalanceRepository {
	db.AutoMigrate(&domain.Balance{})

	return &BalanceRepository{
		db: db,
	}
}

// Function to create balance in database
func (r *BalanceRepository) Create(balance *domain.Balance) (*domain.Balance, error) {
	err := r.db.Create(balance).Error
	if err != nil {
		log.Println("[balances] [repository] error creating balance, err: ", err.Error())
		return nil, err
	}

	return balance, nil
}

// Function to get balance by user id
func (r *BalanceRepository) GetByUserID(id uint) (*domain.Balance, error) {
	var balance domain.Balance

	err := r.db.Where("user_id = ?", id).First(&balance).Error
	if err != nil {
		log.Println("[balances] [repository] error getting balance by user id, err: ", err.Error())
		return nil, err
	}

	return &balance, nil
}

// Function to update balance in database
func (r *BalanceRepository) Update(id uint, balance *domain.Balance) (*domain.Balance, error) {
	var oldBalance domain.Balance

	err := r.db.Where("id = ?", id).First(&oldBalance).Error
	if err != nil {
		log.Println("[balances] [repository] error getting balance by id, err: ", err.Error())
		return nil, err
	}

	err = r.db.Model(&oldBalance).Updates(balance).Error
	if err != nil {
		log.Println("[balances] [repository] error updating balance, err: ", err.Error())
		return nil, err
	}

	return balance, nil
}
