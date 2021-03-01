package peatio

type GetAccountBalanceParams struct {
	UID      string `json:"uid"`
	Currency string `json:"currency"`
}

type CreateWithdrawParams struct {
	UID           string  `json:"uid"`
	TID           string  `json:"tid,omitempty"`
	RID           string  `json:"rid,omitempty"`
	BeneficiaryID string  `json:"beneficiary_id,omitempty"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	Note          string  `json:"note,omitempty"`
	Action        string  `json:"action,omitempty"`
	TransferType  string  `json:"transfer_type,omitempty"`
}

type GenerateDepositAddressParams struct {
	UID      string `json:"uid"`
	Currency string `json:"currency"`
	Remote   bool   `json:"remote,omitempty"`
}

type CreateDepositParams struct {
	UID          string  `json:"uid"`
	TID          string  `json:"tid,omitempty"`
	Currency     string  `json:"currency"`
	Amount       float64 `json:"amount"`
	State        string  `json:"state,omitempty"`
	TransferType string  `json:"transfer_type,omitempty"`
}

type GetDepositsParams struct {
	UID      string `json:"uid,omitempty"`
	FromID   int64  `json:"from_id,omitempty"`
	Currency string `json:"currency,omitempty"`
	Page     int64  `json:"page,omitempty"`
	Limit    int64  `json:"limit,omitempty"`
	State    string `json:"state,omitempty"`
}
