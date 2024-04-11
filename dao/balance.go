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

type BalanceDaoInterface interface {
	List(dto dto.ListBalanceDto) ([]models.Balance, error)
	Create(dto dto.CreateBalanceDto) (*models.Balance, error)
	Get(dto dto.GetBalanceDto) (*models.Balance, error)
	Update(dto dto.UpdateBalanceDto) (*models.Balance, error)
	Delete(dto.DeleteBalanceDto) error
}

type BalanceDao struct {
	db *gorm.DB
}

// List implements BalanceDaoInterface.
func (l *BalanceDao) List(dto dto.ListBalanceDto) (balances []models.Balance, err error) {
	l.db.
		Order("created_at DESC").
		Find(&balances).
		Offset(dto.Offset).
		Limit(dto.Limit)

	return balances, nil
}

// Create implements BalanceDaoInterface.
func (l *BalanceDao) Create(dto dto.CreateBalanceDto) (*models.Balance, error) {
	meta, _ := json.Marshal(dto.Meta)
	balance := models.Balance{
		Ledger: models.Ledger{
			Id: dto.LedgerId,
		},
		TotalBalance:          *dto.TotalBalance,
		InflightBalance:       *dto.InflightBalance,
		CreditBalance:         *dto.CreditBalance,
		DebitBalance:          *dto.DebitBalance,
		InflightCreditBalance: *dto.InflightCreditBalance,
		InflightDebitBalance:  *dto.InflightDebitBalance,
		InflighExpiresAt:      *dto.InflighExpiresAt,
		CurrencyMultiplier:    *dto.CurrencyMultiplier,
		Currency:              dto.Currency,
		Version:               *dto.Version,
		Indicator:             *dto.Indicator,
		Meta:                  meta,
	}

	l.db.Create(&balance)

	return &balance, nil
}

// Get implements BalanceDaoInterface.
func (l *BalanceDao) Get(dto dto.GetBalanceDto) (balance *models.Balance, err error) {
	balanceDoesNotExistErr := l.db.
		Where("id = ?", dto.Id).
		First(&balance)
	balanceDoesNotExist := errors.Is(balanceDoesNotExistErr.Error, gorm.ErrRecordNotFound)
	if balanceDoesNotExist {
		return nil, errors.New("balance does not exist")
	}

	return balance, nil
}

// Update implements BalanceDaoInterface.
func (l *BalanceDao) Update(dto dto.UpdateBalanceDto) (*models.Balance, error) {
	balanceDoesNotExistErr := l.db.
		Where("id = ?", dto.Id).
		First(&models.Balance{})
	balanceDoesNotExist := errors.Is(balanceDoesNotExistErr.Error, gorm.ErrRecordNotFound)
	if balanceDoesNotExist {
		return nil, errors.New("balance does not exist")
	}

	var balance models.Balance
	l.db.Model(&balance).Clauses(clause.Returning{}).Where("id = ?", dto.Id).Updates(dto)

	return &balance, nil
}

// Delete implements BalanceDaoInterface.
func (l *BalanceDao) Delete(dto dto.DeleteBalanceDto) error {
	balanceDoesNotExistErr := l.db.Where("id = ?", dto.Id).First(&models.Balance{})
	balanceDoesNotExist := errors.Is(balanceDoesNotExistErr.Error, gorm.ErrRecordNotFound)
	if balanceDoesNotExist {
		return errors.New("balance does not exist")
	}

	l.db.Where("id = ?", dto.Id).Delete(&models.Balance{})

	return nil
}

func NewBalanceDao() BalanceDaoInterface {
	db, err := core.NewDatabaseConnection()
	if err != nil {
		log.Panicln(err)
	}

	return &BalanceDao{db: db}
}
