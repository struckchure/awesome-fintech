package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Balance struct {
	Id        string         `json:"id" gorm:"primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	LedgerId              string         `json:"ledgerId"`
	Ledger                Ledger         `json:"ledger,omitempty" gorm:"foreignKey:LedgerId"`
	TotalBalance          int64          `json:"totalBalance" gorm:"default:0"`
	InflightBalance       int64          `json:"inflightBalance" gorm:"default:0"`
	CreditBalance         int64          `json:"creditBalance" gorm:"default:0"`
	DebitBalance          int64          `json:"debitBalance" gorm:"default:0"`
	InflightCreditBalance int64          `json:"inflightCreditBalance" gorm:"default:0"`
	InflightDebitBalance  int64          `json:"inflightDebitBalance" gorm:"default:0"`
	InflighExpiresAt      time.Time      `json:"inflightExpiresAt"`
	CurrencyMultiplier    int64          `json:"preceision"`
	Currency              string         `json:"currency"`
	Version               int64          `json:"version"`
	Indicator             string         `json:"indicator"`
	Meta                  datatypes.JSON `json:"meta"`
}
