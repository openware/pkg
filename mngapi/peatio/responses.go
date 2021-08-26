package peatio

type Withdraw struct {
	TID            string `json:"tid"`
	UID            string `json:"uid"`
	Currency       string `json:"currency"`
	BlockchainKey  string `json:"blockchain_key"`
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
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Homepage    string               `json:"homepage"`
	Price       string               `json:"price"`
	Status      string               `json:"status"`
	Type        string               `json:"type"`
	Precision   uint64               `json:"precision"`
	Position    uint64               `json:"position"`
	IconURL     string               `json:"icon_url"`
	Code        string               `json:"code"`
	Networks    []BlockchainCurrency `json:"networks"`
}

type BlockchainCurrency struct {
	ID                  string                 `json:"id"`
	CurrencyID          string                 `json:"currency_id"`
	BlockchainKey       string                 `json:"blockchain_key"`
	ParentID            string                 `json:"parent_id"`
	Status              string                 `json:"status"`
	DepositEnabled      bool                   `json:"deposit_enabled"`
	WithdrawEnabled     bool                   `json:"withdrawal_enabled"`
	DepositFee          string                 `json:"deposit_fee"`
	MinDepositAmount    string                 `json:"min_deposit_amount"`
	WithdrawFee         string                 `json:"withdraw_fee"`
	MinWithdrawAmount   string                 `json:"min_withdraw_amount"`
	BaseFactor          uint64                 `json:"base_factor"`
	MinCollectionAmount string                 `json:"min_collection_amount"`
	Options             map[string]interface{} `json:"options"`
}

type Balance struct {
	UID     string `json:"uid"`
	Balance string `json:"balance"`
	Locked  string `json:"locked"`
}

type PaymentAddress struct {
	UID           string   `json:"uid"`
	Address       string   `json:"address"`
	BlockchainKey string   `json:"blockchain_key"`
	Currencies    []string `json:"currencies"`
	State         string   `json:"state"`
	Remote        bool     `json:"remote"`
}

type Deposit struct {
	ID                      uint64  `json:"id"`
	TID                     string  `json:"tid"`
	BlockchainKey           string  `json:"blockchain_key"`
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
