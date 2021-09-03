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
	AdditionalFee bool    `json:"additional_fee,omitempty"`
}

type GenerateDepositAddressParams struct {
	UID           string `json:"uid"`
	Currency      string `json:"currency"`
	BlockchainKey string `json:"blockchain_key,omitempty"`
	Remote        bool   `json:"remote,omitempty"`
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
	UID           string `json:"uid,omitempty"`
	FromID        int64  `json:"from_id,omitempty"`
	Currency      string `json:"currency,omitempty"`
	BlockchainKey string `json:"blockchain_key,omitempty"`
	Page          int64  `json:"page,omitempty"`
	Limit         int64  `json:"limit,omitempty"`
	State         string `json:"state,omitempty"`
}

type GetEngineParams struct {
	Name string `json:"name"`
}

type CreateEngineParams struct {
	Name   string `json:"name"`
	Driver string `json:"driver"`
	UID    string `json:"uid"`
	URL    string `json:"url"`
	State  int    `json:"state"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type UpdateEngineParams struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Driver string `json:"driver"`
	UID    string `json:"uid"`
	URL    string `json:"url"`
	State  int    `json:"state"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type CreateMarketParams struct {
	BaseCurrency    string `json:"base_currency"`
	QuoteCurrency   string `json:"quote_currency"`
	State           string `json:"state"`
	EngineName      string `json:"engine_name"`
	AmountPrecision int64  `json:"amount_precision"`
	PricePrecision  int64  `json:"price_precision"`
	MinPrice        string `json:"min_price"`
	MaxPrice        string `json:"max_price"`
	MinAmount       string `json:"min_amount"`
	Position        int64  `json:"position"`
}

type UpdateMarketParams struct {
	ID        string `json:"id"`
	EngineID  string `json:"engine_id,omitempty"`
	MinPrice  string `json:"min_price,omitempty"`
	MaxPrice  string `json:"max_price,omitempty"`
	MinAmount string `json:"min_amount,omitempty"`
	AmountPrecision int64  `json:"amount_precision,omitempty"`
	PricePrecision  int64  `json:"price_precision,omitempty"`
}

type CurrenciesListParams struct {
	Type string `json:"type,omitempty"`
}

type CreateBlockchainCurrencyParams struct {
	CurrencyID          string                 `json:"currency_id"`
	BlockchainKey       string                 `json:"blockchain_key,omitempty"`
	BaseFactor          int64                  `json:"base_factor,omitempty"`
	ParentID            string                 `json:"parent_id,omitempty"`
	DepositFee          string                 `json:"deposit_fee,omitempty"`
	MinDepositAmount    string                 `json:"min_deposit_amount,omitempty"`
	MinCollectionAmount string                 `json:"min_collection_amount,omitempty"`
	WithdrawFee         string                 `json:"withdraw_fee,omitempty"`
	MinWithdrawAmount   string                 `json:"min_withdraw_amount,omitempty"`
	DepositEnabled      bool                   `json:"deposit_enabled,omitempty"`
	WithdrawEnabled     bool                   `json:"withdrawal_enabled,omitempty"`
	Status              string                 `json:"status,omitempty"`
	Options             map[string]interface{} `json:"options"`
}

type CreateCurrencyParams struct {
	Code        string `json:"code"`
	Type        string `json:"type,omitempty"`
	Position    int64  `json:"position,omitempty"`
	Name        string `json:"name,omitempty"`
	Precision   int64  `json:"precision,omitempty"`
	Price       string `json:"price,omitempty"`
	Status      string `json:"status,omitempty"`
	IconURL     string `json:"icon_url,omitempty"`
	Description string `json:"description,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
}

type UpdateCurrencyParams struct {
	ID        string `json:"id"`
	Name      string `json:"name,omitempty"`
	Position  int64  `json:"position,omitempty"`
	Status    string `json:"status,omitempty"`
	Precision int64  `json:"precision,omitempty"`
	IconURL   string `json:"icon_url,omitempty"`
}

type UpdateBlockchainCurrencyParams struct {
	ID                  string                 `json:"id"`
	DepositFee          string                 `json:"deposit_fee,omitempty"`
	MinDepositAmount    string                 `json:"min_deposit_amount,omitempty"`
	MinCollectionAmount string                 `json:"min_collection_amount,omitempty"`
	WithdrawFee         string                 `json:"withdraw_fee,omitempty"`
	MinWithdrawAmount   string                 `json:"min_withdraw_amount,omitempty"`
	DepositEnabled      bool                   `json:"deposit_enabled,omitempty"`
	WithdrawEnabled     bool                   `json:"withdrawal_enabled,omitempty"`
	Status              string                 `json:"status,omitempty"`
	Options             map[string]interface{} `json:"options"`
}

type CreateMemberParams struct {
	Email string `json:"email"`
	UID   string `json:"uid"`
	Level int    `json:"level"`
	Role  string `json:"role"`
	State string `json:"state"`
	Group string `json:"group"`
}

type Settings struct {
	URI    string `json:"uri,omitempty"`
	Secret string `json:"secret,omitempty"`
}
type CreateWalletParams struct {
	BlockchainKey string   `json:"blockchain_key"`
	Name          string   `json:"name"`
	Kind          string   `json:"kind"`
	Gateway       string   `json:"gateway"`
	Address       string   `json:"address"`
	Currencies    []string `json:"currencies,omitempty"`
	Settings      Settings `json:"settings,omitempty"`
	MaxBalance    string   `json:"max_balance,omitempty"`
	Status        string   `json:"status,omitempty"`
}

type UpdateWalletParams struct {
	ID            string   `json:"id"`
	BlockchainKey string   `json:"blockchain_key,omitempty"`
	Name          string   `json:"name,omitempty"`
	Address       string   `json:"address,omitempty"`
	Gateway       string   `json:"gateway,omitempty"`
	Kind          string   `json:"kind,omitempty"`
	Currencies    []string `json:"currencies,omitempty"`
	Settings      Settings `json:"settings,omitempty"`
	MaxBalance    string   `json:"max_balance,omitempty"`
	Status        string   `json:"status,omitempty"`
}
