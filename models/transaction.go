package models

import (
	"encoding/json"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Transaction struct {
	Id        string         `json:"id" gorm:"primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	Source      string         `json:"source"`
	Destination string         `json:"destination"`
	Reference   string         `json:"reference"`
	Amount      int64          `json:"amount"`
	Currency    string         `json:"currency"`
	Status      string         `json:"status" gorm:"default:Pending"`
	Meta        datatypes.JSON `json:"meta,omitempty"`
}

func (transaction *Transaction) ToJSON() ([]byte, error) {
	return json.Marshal(transaction)
}
