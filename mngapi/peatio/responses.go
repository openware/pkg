package peatio

type Withdraw struct {
	TID            string  `json:"tid"`
	UID            string  `json:"uid"`
	Currency       string  `json:"currency"`
	Note           string  `json:"note"`
	Type           string  `json:"type"`
	Amount         string  `json:"amount"`
	Fee            string  `json:"fee"`
	RID            string  `json:"rid"`
	State          string  `json:"state"`
	CreatedAt      string  `json:"created_at"`
	BlockchainTxID *string `json:"blockchain_txid"`
	TransferType   string  `json:"transfer_type"`
}

type Currency struct {
	ID                  string                 `json:"id"`
	Name                string                 `json:"name"`
	Description         string                 `json:"description"`
	Homepage            string                 `json:"homepage"`
	Price               string                 `json:"price"`
	ParentID            string                 `json:"parent_id"`
	ExplorerTransaction string                 `json:"explorer_transaction"`
	ExplorerAddress     string                 `json:"explorer_address"`
	Type                string                 `json:"type"`
	DepositEnabled      bool                   `json:"deposit_enabled"`
	WithdrawEnabled     bool                   `json:"withdrawal_enabled"`
	DepositFee          string                 `json:"deposit_fee"`
	MinDepositAmount    string                 `json:"min_deposit_amount"`
	WithdrawFee         string                 `json:"withdraw_fee"`
	MinWithdrawAmount   string                 `json:"min_withdraw_amount"`
	WithdrawLimit24h    string                 `json:"withdraw_limit_24h"`
	WithdrawLimit72h    string                 `json:"withdraw_limit_72h"`
	BaseFactor          uint64                 `json:"base_factor"`
	Precision           uint64                 `json:"precision"`
	Position            uint64                 `json:"position"`
	IconURL             string                 `json:"icon_url"`
	MinConfirmations    uint64                 `json:"min_confirmations"`
	Code                string                 `json:"code"`
	MinCollectionAmount string                 `json:"min_collection_amount"`
	Visible             bool                   `json:"visible"`
	SubUnits            uint64                 `json:"subunits"`
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

type Deposit struct {
	ID                      uint64  `json:"id"`
	TID                     string  `json:"tid"`
	Currency                string  `json:"currency"`
	Address                 string  `json:"address"`
	UID                     string  `json:"uid"`
	Type                    string  `json:"type"`
	Amount                  string  `json:"amount"`
	State                   string  `json:"state"`
	CreatedAt               string  `json:"created_at"`
	CompletedAt             *string `json:"completed_at"`
	BlockchainTxID          string  `json:"blockchain_txid,omitempty"`
	BlockchainConfirmations string  `json:"blockchain_confirmations,omitempty"`
	TransferType            string  `json:"transfer_type"`
}

type Engine struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Driver string `json:"driver"`
	UID    string `json:"uid"`
	URL    string `json:"url"`
	State  string `json:"state"`
}

type Market struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	BaseUnit        string `json:"base_unit"`
	QuoteUnit       string `json:"quote_unit"`
	MinPrice        string `json:"min_price"`
	MaxPrice        string `json:"max_price"`
	MinAmount       string `json:"min_amount"`
	AmountPrecision int    `json:"amount_precision"`
	PricePrecision  int    `json:"price_precision"`
	State           string `json:"state"`
	Position        int    `json:"position"`
	EngineID        int    `json:"engine_id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type Member struct {
	UID   string `json:"uid"`
	Email string `json:"email"`
	Level int    `json:"level"`
	Role  string `json:"role"`
	Group string `json:"group"`
	State string `json:"state"`
}

type Wallet struct {
	ID            int                    `json:"id"`
	Name          string                 `json:"name"`
	Kind          string                 `json:"kind"`
	Currencies    []string               `json:"currencies"`
	Address       string                 `json:"address"`
	Gateway       string                 `json:"gateway"`
	MaxBalance    string                 `json:"max_balance"`
	Balance       map[string]interface{} `json:"balance"`
	BlockchainKey string                 `json:"blockchain_key"`
	Status        string                 `json:"status"`
}
