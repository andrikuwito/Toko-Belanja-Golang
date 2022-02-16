package repository

import (
	"tokobelanja-golang/model/entity"

	"gorm.io/gorm"
)

type TransactionHistoryRepository interface {
	GetTransactionByIDUser(ID int) ([]entity.TransactionHistory, error)
	GetAllTransaction() ([]entity.TransactionHistory, error)
	Save(input entity.TransactionHistory) (entity.TransactionHistory, error)
}

type transactionHistoryRepository struct {
	db *gorm.DB
}

func NewTransactionHistoryRepository(db *gorm.DB) *transactionHistoryRepository {
	return &transactionHistoryRepository{db}
}

func (r *transactionHistoryRepository) GetTransactionByIDUser(ID int) ([]entity.TransactionHistory, error) {
	allTransactions := []entity.TransactionHistory{}
	err := r.db.Where("user_id", ID).Find(&allTransactions).Error

	if err != nil {
		return []entity.TransactionHistory{}, err
	}

	return allTransactions, nil
}

func (r *transactionHistoryRepository) GetAllTransaction() ([]entity.TransactionHistory, error) {
	allTransactions := []entity.TransactionHistory{}
	err := r.db.Find(&allTransactions).Error

	if err != nil {
		return []entity.TransactionHistory{}, err
	}

	return allTransactions, nil
}

func (r *transactionHistoryRepository) Save(input entity.TransactionHistory) (entity.TransactionHistory, error) {
	err := r.db.Create(&input).Error

	if err != nil {
		return entity.TransactionHistory{}, err
	}

	return input, nil
}
