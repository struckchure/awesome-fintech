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

	Ledger         Ledger         `json:"ledger,omitempty" gorm:"foreignKey:LedgerId"`
	LedgerId       string         `json:"ledgerId"`
	AllowOverdraft bool           `json:"allowOverdraft" gorm:"default:false"`
	TotalBalance   int64          `json:"totalBalance" gorm:"default:0"`
	Currency       string         `json:"currency"`
	Meta           datatypes.JSON `json:"meta"`
}
