package models

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Transaction struct {
	Id        string         `json:"id" gorm:"primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	AllowOverdraft         bool           `json:"allowOverdraft" gorm:"default:false"`
	Source                 string         `json:"source"`
	Destination            string         `json:"destination"`
	Reference              string         `json:"reference"`
	Amount                 int64          `json:"amount"`
	Currency               string         `json:"currency"`
	Description            string         `json:"description"`
	Status                 string         `json:"status" gorm:"default:Pending"`
	ScheduledFor           time.Time      `json:"scheduledFor,omitempty"`
	RiskToleranceThreshold float64        `json:"riskToleranceThreshold"`
	RiskScore              float64        `json:"riskScore"`
	SkipBalanceUpdate      bool           `json:"-"`
	Hash                   string         `json:"hash"`
	Meta                   datatypes.JSON `json:"meta,omitempty"`
	GroupIds               pq.StringArray `json:"groupIds" gorm:"type:text[]"`
}

func (transaction *Transaction) ToJSON() ([]byte, error) {
	return json.Marshal(transaction)
}
