package models

import (
	"time"

	"github.com/nusa-exchange/finex/config"
	"github.com/shopspring/decimal"
)

type Deposit struct {
	ID            int64           `json:"id"`
	MemberID      int64           `json:"member_id"`
	CurrencyID    string          `json:"currency_id"`
	Amount        decimal.Decimal `json:"amount"`
	Fee           decimal.Decimal `json:"fee"`
	Txid          string          `json:"txid"`
	AasmState     string          `json:"aasm_state"`
	CreatedAT     time.Time       `json:"created_at"`
	UpdatedAT     time.Time       `json:"updated_at"`
	CompletedAT   time.Time       `json:"completed_at"`
	Type          string          `json:"type"`
	Txout         int8            `json:"txout"`
	Tid           string          `json:"tid"`
	Address       string          `json:"address"`
	BlockNumber   int64           `json:"block_number"`
	Spread        string          `json:"spread"`
	FromAddresses string          `json:"from_addresses"`
	TransferType  int8            `json:"transfer_type"`
	Error         []byte          `json:"error"`
	BlockchainKey string          `json:"blockchain_key"`
}

func (d *Deposit) Members() *Member {
	var member *Member
	config.DataBase.Where("id = ?", d.MemberID).First(&member)

	return member
}

func (d *Deposit) Currency() *Currency {
	var currency *Currency
	config.DataBase.Where("id = ?", d.CurrencyID).First(&currency)

	return currency
}

func (d *Deposit) Blockchain() *Blockchains {
	var blockchain *Blockchains
	config.DataBase.Where("key = ?", d.BlockchainKey).First(&blockchain)

	return blockchain
}
