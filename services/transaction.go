package services

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"

	"awesome.fintech.org/core/clients"
	"awesome.fintech.org/core/constants"
	"awesome.fintech.org/dao"
	"awesome.fintech.org/dto"
	"awesome.fintech.org/models"
)

type TransactionService struct {
	rabbitmq *clients.RabbitMQ

	transactionDao dao.TransactionDaoInterface
}

func (s *TransactionService) List(args dto.ListTransactionDto) ([]models.Transaction, error) {
	return s.transactionDao.List(args)
}

func (s *TransactionService) Record(args dto.CreateTransactionDto) (*models.Transaction, error) {
	args.Reference = uuid.NewString()

	transaction, err := s.transactionDao.Create(args)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(transaction)
	if err != nil {
		log.Printf("Error marshalling transaction: %s", err)
	}

	s.rabbitmq.Publish(clients.PublishArgs{
		Queue:   constants.TRANSACTIONS_QUEUE,
		Content: string(content),
	})

	return transaction, nil
}

func (s *TransactionService) HandleRecord(args dto.CreateTransactionDto) {}

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
		Source:      transaction.Destination, // switch `Source` for `Destination`
		Destination: transaction.Source,      // switch `Destination` for `Source`
		Reference:   transaction.Reference,
		Amount:      transaction.Amount,
		Currency:    transaction.Currency,
		Status:      transaction.Status,
		Meta:        meta,
	})
	if err != nil {
		return nil, err
	}

	return refundTransaction, nil
}

func NewTransactionService(rabbitmq *clients.RabbitMQ, transactionDao dao.TransactionDaoInterface) *TransactionService {
	return &TransactionService{
		rabbitmq: rabbitmq,

		transactionDao: transactionDao,
	}
}
