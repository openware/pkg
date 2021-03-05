package peatio

type Withdraw struct {
	TID            string `json:"tid"`
	UID            string `json:"uid"`
	Currency       string `json:"currency"`
	Note           string `json:"note"`
	Type           string `json:"type"`
	Amount         string `json:"amount"`
	Fee            string `json:"fee"`
	RID            string `json:"rid"`
	State          string `json:"state"`
	CreatedAt      string `json:"created_at"`
	BlockchainTxID string `json:"blockchain_txid"`
	TransferType   string `json:"transfer_type"`
}

type Currency struct {
	ID                  string                 `json:"id"`
	Name                string                 `json:"name"`
	Description         string                 `json:"descritpion"`
	Homepage            string                 `json:"homepage"`
	Price               string                 `json:"price"`
	ExplorerTransaction string                 `json:"explorer_transaction"`
	ExplorerAddress     string                 `json:"explorer_address"`
	Type                string                 `json:"type"`
	DepositEnabled      string                 `json:"deposit_enabled"`
	WithdrawEnabled     string                 `json:"withdrawal_enabled"`
	DepositFee          string                 `json:"deposit_fee"`
	MinDepositAmount    string                 `json:"min_deposit_amount"`
	WithdrawFee         string                 `json:"withdraw_fee"`
	MinWithdrawAmount   string                 `json:"min_withdraw_amount"`
	WithdrawLimit24h    string                 `json:"withdraw_limit_24h"`
	WithdrawLimit72h    string                 `json:"withdraw_limit_72h"`
	BaseFactor          string                 `json:"base_factor"`
	Precision           string                 `json:"precision"`
	Position            int                    `json:"position"`
	IconURL             string                 `json:"icon_url"`
	MinConfirmations    string                 `json:"min_confirmations"`
	Code                string                 `json:"code"`
	MinCollectionAmount string                 `json:"min_collection_amount"`
	Visible             string                 `json:"visible"`
	SubUnits            int                    `json:"subunits"`
	Options             map[string]interface{} `json:"options"`
	CreatedAt           string                 `json:"created_at"`
	UpdatedAt           string                 `json:"updated_at"`
}

type Balance struct {
	UID     string `json:"uid"`
	Balance string `json:"balance"`
	Locked  string `json:"locked"`
}

type PaymentAddress struct {
	UID        string   `json:"uid"`
	Address    string   `json:"address"`
	Currencies []string `json:"currencies"`
	State      string   `json:"state"`
	Remote     bool     `json:"remote"`
}
