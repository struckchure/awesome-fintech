package dto

type ListLedgerDto struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type CreateLedgerDto struct {
	Name string                 `json:"name"`
	Meta map[string]interface{} `json:"meta"`
}

type GetLedgerDto struct {
	Id string `json:"id"`
}

type UpdateLedgerDto struct {
	Id   string                  `json:"id"`
	Name *string                 `json:"name,omitempty"`
	Meta *map[string]interface{} `json:"meta,omitempty"`
}

type DeleteLedgerDto struct {
	Id string `json:"id"`
}
