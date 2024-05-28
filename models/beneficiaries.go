package models

import (
	"time"

	"github.com/nusa-exchange/finex/config"
)

type Beneficiary struct {
	ID            int64     `json:"id"`
	MemberID      int64     `json:"member_id"`
	CurrencyID    int64     `json:"currency_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Pin           int32     `json:"pin"`
	State         int8      `json:"state"`
	CreatedT      time.Time `json:"created_at"`
	UpdatedAT     time.Time `json:"updated_at"`
	SentAT        time.Time `json:"sent_at"`
	DataEncrypted string    `json:"data_encrypted"`
	BlockchainKey string    `json:"blockchain_key"`
	ExpireAT      time.Time `json:"expire_at"`
}

func (d *Beneficiary) Members() *Member {
	var member *Member
	config.DataBase.Where("id = ?", d.MemberID).First(&member)

	return member
}

func (d *Beneficiary) Currency() *Currency {
	var currency *Currency
	config.DataBase.Where("id = ?", d.CurrencyID).First(&currency)

	return currency
}

func (d *Beneficiary) Blockchain() *Blockchains {
	var blockchain *Blockchains
	config.DataBase.Where("key = ?", d.BlockchainKey).First(&blockchain)

	return blockchain
}

func (d *Beneficiary) Data() *Blockchains {
	var blockchain *Blockchains
	config.DataBase.Where("key = ?", d.BlockchainKey).First(&blockchain)

	return blockchain
}
