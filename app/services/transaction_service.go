package services

import (
	"myapp/app/repositories"
	"myapp/models"
)

type TransactionService interface {
	GetAll() ([]models.Transaction, error)
	Create(tx *models.Transaction) error
}

type transactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(
	repo repositories.TransactionRepository,
) TransactionService {
	return &transactionService{
		repo: repo,
	}
}

func (s *transactionService) GetAll() (
	[]models.Transaction,
	error,
) {
	return s.repo.FindAll()
}

func (s *transactionService) Create(
	tx *models.Transaction,
) error {
	tx.Status = "pending"

	return s.repo.Create(tx)
}
