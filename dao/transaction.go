package dao

import (
	"encoding/json"
	"errors"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"awesome.fintech.org/core"
	"awesome.fintech.org/dto"
	"awesome.fintech.org/models"
)

type TransactionDaoInterface interface {
	List(dto dto.ListTransactionDto) ([]models.Transaction, error)
	Create(dto dto.CreateTransactionDto) (*models.Transaction, error)
	Get(dto dto.GetTransactionDto) (*models.Transaction, error)
	Update(dto dto.UpdateTransactionDto) (*models.Transaction, error)
	Delete(dto.DeleteTransactionDto) error
}

type TransactionDao struct {
	db *gorm.DB
}

// List implements TransactionDaoInterface.
func (l *TransactionDao) List(dto dto.ListTransactionDto) (transactions []models.Transaction, err error) {
	l.db.
		Order("created_at DESC").
		Find(&transactions).
		Offset(dto.Offset).
		Limit(dto.Limit)

	return transactions, nil
}

// Create implements TransactionDaoInterface.
func (l *TransactionDao) Create(dto dto.CreateTransactionDto) (*models.Transaction, error) {
	meta, _ := json.Marshal(dto.Meta)
	transaction := models.Transaction{
		AllowOverdraft:         *dto.AllowOverdraft,
		Source:                 dto.Source,
		Destination:            dto.Destination,
		Reference:              dto.Reference,
		Amount:                 dto.Amount,
		Currency:               dto.Currency,
		Description:            *dto.Description,
		Status:                 *dto.Status,
		ScheduledFor:           *dto.ScheduledFor,
		RiskToleranceThreshold: *dto.RiskToleranceThreshold,
		RiskScore:              *dto.RiskScore,
		Hash:                   *dto.Hash,
		GroupIds:               dto.GroupIds,
		Meta:                   meta,
	}

	l.db.Create(&transaction)

	return &transaction, nil
}

// Get implements TransactionDaoInterface.
func (l *TransactionDao) Get(dto dto.GetTransactionDto) (transaction *models.Transaction, err error) {
	transactionDoesNotExistErr := l.db.
		Where("id = ?", dto.Id).
		First(&transaction)
	transactionDoesNotExist := errors.Is(transactionDoesNotExistErr.Error, gorm.ErrRecordNotFound)
	if transactionDoesNotExist {
		return nil, errors.New("transaction does not exist")
	}

	return transaction, nil
}

// Update implements TransactionDaoInterface.
func (l *TransactionDao) Update(dto dto.UpdateTransactionDto) (*models.Transaction, error) {
	transactionDoesNotExistErr := l.db.
		Where("id = ?", dto.Id).
		First(&models.Transaction{})
	transactionDoesNotExist := errors.Is(transactionDoesNotExistErr.Error, gorm.ErrRecordNotFound)
	if transactionDoesNotExist {
		return nil, errors.New("transaction does not exist")
	}

	var transaction models.Transaction
	l.db.Model(&transaction).Clauses(clause.Returning{}).Where("id = ?", dto.Id).Updates(dto)

	return &transaction, nil
}

// Delete implements TransactionDaoInterface.
func (l *TransactionDao) Delete(dto dto.DeleteTransactionDto) error {
	transactionDoesNotExistErr := l.db.Where("id = ?", dto.Id).First(&models.Transaction{})
	transactionDoesNotExist := errors.Is(transactionDoesNotExistErr.Error, gorm.ErrRecordNotFound)
	if transactionDoesNotExist {
		return errors.New("transaction does not exist")
	}

	l.db.Where("id = ?", dto.Id).Delete(&models.Transaction{})

	return nil
}

func NewTransactionDao() TransactionDaoInterface {
	db, err := core.NewDatabaseConnection()
	if err != nil {
		log.Panicln(err)
	}

	return &TransactionDao{db: db}
}
