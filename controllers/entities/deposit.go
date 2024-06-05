package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type DepositEntities struct {
	ID            int64           `json:"id"`
	Currency      string          `json:"currency"`
	BlockchainKey string          `json:"blockchain_key"`
	Protocol      string          `json:"protocol"`
	Warning       string          `json:"warning"`
	Amount        decimal.Decimal `json:"amount"`
	Fee           decimal.Decimal `json:"fee"`
	FromAddresses string          `json:"from_addresses"`
	Txid          string          `json:"txid"`
	Confirmations int64           `json:"confirmations"`
	State         string          `json:"state"`
	TransferType  string          `json:"transfer_type"`
	Error         string          `json:"error"`
	Tid           string          `json:"tid"`
	CreatedAT     time.Time       `json:"created_at"`
	CompletedAT   time.Time       `json:"completed_at"`
}
