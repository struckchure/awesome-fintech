package services

import (
	"awesome.fintech.org/dao"
	"awesome.fintech.org/dto"
	"awesome.fintech.org/models"
)

type LedgerService struct {
	ledgerDao dao.LedgerDaoInterface
}

func (s *LedgerService) List(dto dto.ListLedgerDto) ([]models.Ledger, error) {
	return s.ledgerDao.List(dto)
}

func (s *LedgerService) Create(dto dto.CreateLedgerDto) (*models.Ledger, error) {
	return s.ledgerDao.Create(dto)
}

func NewLedgerService() *LedgerService {
	return &LedgerService{
		ledgerDao: dao.NewLedgerDao(),
	}
}
