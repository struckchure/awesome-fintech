package dto

type ListBalanceDto struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type CreateBalanceDto struct {
	LedgerId       string                  `json:"ledgerId"`
	TotalBalance   *int64                  `json:"totalBalance"`
	AllowOverdraft bool                    `json:"allowOverdraft"`
	Currency       string                  `json:"currency"`
	Meta           *map[string]interface{} `json:"meta"`
}

type GetBalanceDto struct {
	Id string `json:"id"`
}

type UpdateBalanceDto struct {
	Id             string                  `json:"id"`
	LedgerId       *string                 `json:"ledgerId"`
	TotalBalance   *int64                  `json:"totalBalance"`
	AllowOverdraft *bool                   `json:"allowOverdraft"`
	Currency       *string                 `json:"currency"`
	Meta           *map[string]interface{} `json:"meta"`
}

type DeleteBalanceDto struct {
	Id string `json:"id"`
}
