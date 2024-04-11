package dto

import "time"

type ListTransactionDto struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type CreateTransactionDto struct {
	AllowOverdraft         *bool                   `json:"allowOverdraft"`
	Source                 string                  `json:"source"`
	Destination            string                  `json:"destination"`
	Reference              string                  `json:"reference"`
	Amount                 int64                   `json:"amount"`
	Currency               string                  `json:"currency"`
	Description            *string                 `json:"description,omitempty"`
	Status                 *string                 `json:"status,omitempty"`
	ScheduledFor           *time.Time              `json:"scheduledFor,omitempty"`
	RiskToleranceThreshold *float64                `json:"riskToleranceThreshold,omitempty"`
	RiskScore              *float64                `json:"riskScore,omitempty"`
	Hash                   *string                 `json:"hash,omitempty"`
	Meta                   *map[string]interface{} `json:"meta,omitempty"`
	GroupIds               []string                `json:"groupIds,omitempty"`
}

type GetTransactionDto struct {
	Id string `json:"id"`
}

type UpdateTransactionDto struct {
	Id                     string                  `json:"id"`
	AllowOverdraft         *bool                   `json:"allowOverdraft,omitempty"`
	Source                 *string                 `json:"source,omitempty"`
	Destination            *string                 `json:"destination,omitempty"`
	Reference              *string                 `json:"reference,omitempty"`
	Amount                 *int64                  `json:"amount,omitempty"`
	Currency               *string                 `json:"currency,omitempty"`
	Description            *string                 `json:"description,omitempty"`
	Status                 *string                 `json:"status,omitempty"`
	ScheduledFor           *time.Time              `json:"scheduledFor,omitempty"`
	RiskToleranceThreshold *float64                `json:"riskToleranceThreshold,omitempty"`
	RiskScore              *float64                `json:"riskScore,omitempty"`
	Hash                   *string                 `json:"hash,omitempty"`
	Meta                   *map[string]interface{} `json:"meta,omitempty"`
	GroupIds               *[]string               `json:"groupIds,omitempty"`
}

type DeleteTransactionDto struct {
	Id string `json:"id"`
}
