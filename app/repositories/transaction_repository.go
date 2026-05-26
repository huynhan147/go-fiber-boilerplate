package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAll() ([]models.Transaction, error)
	FindByID(id uint) (*models.Transaction, error)
	Create(tx *models.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(
	db *gorm.DB,
) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) FindAll() (
	[]models.Transaction,
	error,
) {
	var transactions []models.Transaction

	err := r.db.Find(&transactions).Error

	return transactions, err
}

func (r *transactionRepository) FindByID(
	id uint,
) (*models.Transaction, error) {

	var transaction models.Transaction

	err := r.db.First(&transaction, id).Error

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *transactionRepository) Create(
	tx *models.Transaction,
) error {
	return r.db.Create(tx).Error
}
