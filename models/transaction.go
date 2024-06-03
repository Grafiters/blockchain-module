package models

import (
	"time"

	"github.com/nusa-exchange/finex/config"
	"github.com/shopspring/decimal"
)

type TransactionStatus string
type TransactionReference string

var (
	TransactinPending TransactionStatus = "pending"
	TransactinSuccess TransactionStatus = "success"
	TransactinFailed  TransactionStatus = "failed"
	TransactinRevert  TransactionStatus = "revert"
)

var (
	ReferenceDeposit  TransactionReference = "Deposit"
	ReferenceWithDraw TransactionReference = "Withdraw"
)

type Transactions struct {
	ID            int64             `json:"id"`
	CurrencyID    string            `json:"currency_id"`
	ReferenceType string            `json:"reference_type"`
	ReferenceID   int64             `json:"reference_id"`
	Txid          string            `json:"txid"`
	FromAddress   string            `json:"from_address"`
	ToAddress     string            `json:"to_address"`
	Amount        decimal.Decimal   `json:"amount"`
	BlockNumber   int64             `json:"block_number"`
	Txout         int8              `json:"txout"`
	Status        TransactionStatus `json:"status"`
	Options       []byte            `json:"options"`
	CreatedAT     time.Time         `json:"created_at"`
	UpdatedAT     time.Time         `json:"updated_at"`
	BlockchainKey string            `json:"blockchain_key"`
	Kind          string            `json:"kind"`
	Fee           decimal.Decimal   `json:"fee"`
	FeeCurrencyID string            `json:"fee_currency_id"`
}

func (t *Transactions) Blockchain() *Blockchains {
	var blockchain *Blockchains
	config.DataBase.Where("key = ?", t.BlockchainKey).First(&blockchain)
	return blockchain
}

func (t *Transactions) Currency() *Currency {
	var currency *Currency
	config.DataBase.Where("ID = ?", t.CurrencyID).First(&currency)
	return currency
}

func (t *Transactions) FeeCurrency() *Currency {
	var currency *Currency
	config.DataBase.Where("ID = ?", t.FeeCurrencyID).First(&currency)
	return currency
}
