package services

import (
	"encoding/json"

	"github.com/google/uuid"

	"awesome.fintech.org/dao"
	"awesome.fintech.org/dto"
	"awesome.fintech.org/models"
)

type TransactionService struct {
	transactionDao dao.TransactionDaoInterface
}

func (s *TransactionService) Record(args dto.CreateTransactionDto) (*models.Transaction, error) {
	args.Reference = uuid.NewString()

	return s.transactionDao.Create(args)
}

func (s *TransactionService) Refund(args dto.GetTransactionDto) (*models.Transaction, error) {
	transaction, err := s.transactionDao.Get(args)
	if err != nil {
		return nil, err
	}

	var meta map[string]interface{}
	err = json.Unmarshal(transaction.Meta, &meta)
	if err != nil {
		return nil, err
	}

	refundTransaction, err := s.Record(dto.CreateTransactionDto{
		AllowOverdraft:         &transaction.AllowOverdraft,
		Source:                 transaction.Destination, // switch `Source` for `Destination`
		Destination:            transaction.Source,      // switch `Destination` for `Source`
		Reference:              transaction.Reference,
		Amount:                 transaction.Amount,
		Currency:               transaction.Currency,
		Description:            &transaction.Description,
		Status:                 &transaction.Status,
		ScheduledFor:           &transaction.ScheduledFor,
		RiskToleranceThreshold: &transaction.RiskToleranceThreshold,
		RiskScore:              &transaction.RiskScore,
		Hash:                   &transaction.Hash,
		Meta:                   &meta,
		GroupIds:               transaction.GroupIds,
	})
	if err != nil {
		return nil, err
	}

	return refundTransaction, nil
}

func NewTransactionService(transactionDao dao.TransactionDaoInterface) *TransactionService {
	return &TransactionService{
		transactionDao: transactionDao,
	}
}
