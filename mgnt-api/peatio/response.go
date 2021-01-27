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
