package dto

import "time"

type ListBalanceDto struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type CreateBalanceDto struct {
	LedgerId              string                  `json:"ledgerId"`
	TotalBalance          *int64                  `json:"totalBalance"`
	InflightBalance       *int64                  `json:"inflightBalance"`
	CreditBalance         *int64                  `json:"creditBalance"`
	DebitBalance          *int64                  `json:"debitBalance"`
	InflightCreditBalance *int64                  `json:"inflightCreditBalance"`
	InflightDebitBalance  *int64                  `json:"inflightDebitBalance"`
	InflighExpiresAt      *time.Time              `json:"inflighExpiresAt"`
	CurrencyMultiplier    *int64                  `json:"currencyMultiplier"`
	Currency              string                  `json:"currency"`
	Version               *int64                  `json:"version"`
	Indicator             *string                 `json:"indicator"`
	Meta                  *map[string]interface{} `json:"meta"`
}

type GetBalanceDto struct {
	Id string `json:"id"`
}

type UpdateBalanceDto struct {
	Id                    string                  `json:"id"`
	LedgerId              *string                 `json:"ledgerId"`
	TotalBalance          *int64                  `json:"totalBalance"`
	InflightBalance       *int64                  `json:"inflightBalance"`
	CreditBalance         *int64                  `json:"creditBalance"`
	DebitBalance          *int64                  `json:"debitBalance"`
	InflightCreditBalance *int64                  `json:"inflightCreditBalance"`
	InflightDebitBalance  *int64                  `json:"inflightDebitBalance"`
	InflighExpiresAt      *time.Time              `json:"inflighExpiresAt"`
	CurrencyMultiplier    *int64                  `json:"currencyMultiplier"`
	Currency              *string                 `json:"currency"`
	Version               *int64                  `json:"version"`
	Indicator             *string                 `json:"indicator"`
	Meta                  *map[string]interface{} `json:"meta"`
}

type DeleteBalanceDto struct {
	Id string `json:"id"`
}
