package dto

type ListTransactionDto struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type CreateTransactionDto struct {
	AllowOverdraft *bool                   `json:"allowOverdraft"`
	Source         string                  `json:"source"`
	Destination    string                  `json:"destination"`
	Reference      string                  `json:"reference"`
	Amount         int64                   `json:"amount"`
	Currency       string                  `json:"currency"`
	Status         *string                 `json:"status,omitempty"`
	Meta           *map[string]interface{} `json:"meta,omitempty"`
}

type GetTransactionDto struct {
	Id string `json:"id"`
}

type UpdateTransactionDto struct {
	Id             string                  `json:"id"`
	AllowOverdraft *bool                   `json:"allowOverdraft,omitempty"`
	Source         *string                 `json:"source,omitempty"`
	Destination    *string                 `json:"destination,omitempty"`
	Reference      *string                 `json:"reference,omitempty"`
	Amount         *int64                  `json:"amount,omitempty"`
	Currency       *string                 `json:"currency,omitempty"`
	Status         *string                 `json:"status,omitempty"`
	Meta           *map[string]interface{} `json:"meta,omitempty"`
}

type DeleteTransactionDto struct {
	Id string `json:"id"`
}
