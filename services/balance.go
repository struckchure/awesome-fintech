package services

import (
	"awesome.fintech.org/dao"
	"awesome.fintech.org/dto"
	"awesome.fintech.org/models"
)

type BalanceService struct {
	balanceDao dao.BalanceDaoInterface
}

func (s *BalanceService) List(dto dto.ListBalanceDto) ([]models.Balance, error) {
	return s.balanceDao.List(dto)
}

func (s *BalanceService) Create(dto dto.CreateBalanceDto) (*models.Balance, error) {
	return s.balanceDao.Create(dto)
}

func (s *BalanceService) Get(dto dto.GetBalanceDto) (*models.Balance, error) {
	return s.balanceDao.Get(dto)
}

func NewBalanceService(balanceDao dao.BalanceDaoInterface) *BalanceService {
	return &BalanceService{
		balanceDao: balanceDao,
	}
}
