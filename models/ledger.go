package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Ledger struct {
	Id        string         `json:"id" gorm:"primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	Name string         `json:"name"`
	Meta datatypes.JSON `json:"meta"`
}
