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
