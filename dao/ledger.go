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

type LedgerDaoInterface interface {
	List(dto dto.ListLedgerDto) ([]models.Ledger, error)
	Create(dto dto.CreateLedgerDto) (*models.Ledger, error)
	Get(dto dto.GetLedgerDto) (*models.Ledger, error)
	Update(dto dto.UpdateLedgerDto) (*models.Ledger, error)
	Delete(dto.DeleteLedgerDto) error
}

type LedgerDao struct {
	db *gorm.DB
}

// List implements LedgerDaoInterface.
func (l *LedgerDao) List(dto dto.ListLedgerDto) (ledgers []models.Ledger, err error) {
	l.db.
		Order("created_at DESC").
		Find(&ledgers).
		Offset(dto.Offset).
		Limit(dto.Limit)

	return ledgers, nil
}

// Create implements LedgerDaoInterface.
func (l *LedgerDao) Create(dto dto.CreateLedgerDto) (*models.Ledger, error) {
	ledgerExistsErr := l.db.
		Where("name = ?", dto.Name).
		First(&models.Ledger{})
	machinePlanExists := !errors.Is(ledgerExistsErr.Error, gorm.ErrRecordNotFound)
	if machinePlanExists {
		return nil, errors.New("plan already exist")
	}

	meta, _ := json.Marshal(dto.Meta)
	ledger := models.Ledger{Name: dto.Name, Meta: meta}

	l.db.Create(&ledger)

	return &ledger, nil
}

// Get implements LedgerDaoInterface.
func (l *LedgerDao) Get(dto dto.GetLedgerDto) (ledger *models.Ledger, err error) {
	ledgerDoesNotExistErr := l.db.
		Where("id = ?", dto.Id).
		First(&ledger)
	ledgerDoesNotExist := errors.Is(ledgerDoesNotExistErr.Error, gorm.ErrRecordNotFound)
	if ledgerDoesNotExist {
		return nil, errors.New("ledger does not exist")
	}

	return ledger, nil
}

// Update implements LedgerDaoInterface.
func (l *LedgerDao) Update(dto dto.UpdateLedgerDto) (*models.Ledger, error) {
	ledgerDoesNotExistErr := l.db.
		Where("id = ?", dto.Id).
		First(&models.Ledger{})
	ledgerDoesNotExist := errors.Is(ledgerDoesNotExistErr.Error, gorm.ErrRecordNotFound)
	if ledgerDoesNotExist {
		return nil, errors.New("ledger does not exist")
	}

	var ledger models.Ledger
	l.db.Model(&ledger).Clauses(clause.Returning{}).Where("id = ?", dto.Id).Updates(dto)

	return &ledger, nil
}

// Delete implements LedgerDaoInterface.
func (l *LedgerDao) Delete(dto dto.DeleteLedgerDto) error {
	ledgerDoesNotExistErr := l.db.Where("id = ?", dto.Id).First(&models.Ledger{})
	ledgerDoesNotExist := errors.Is(ledgerDoesNotExistErr.Error, gorm.ErrRecordNotFound)
	if ledgerDoesNotExist {
		return errors.New("ledger does not exist")
	}

	l.db.Where("id = ?", dto.Id).Delete(&models.Ledger{})

	return nil
}

func NewLedgerDao() LedgerDaoInterface {
	db, err := core.NewDatabaseConnection()
	if err != nil {
		log.Panicln(err)
	}

	return &LedgerDao{db: db}
}
