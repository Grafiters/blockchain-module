package models

import (
	"time"

	"github.com/nusa-exchange/finex/config"
)

type WhitelistedAddresses struct {
	ID            string    `json:"id"`
	Description   string    `json:"description"`
	Address       string    `json:"address"`
	State         string    `json:"state"`
	BlockchainKey string    `json:"blockchain_key"`
	CreatedAT     time.Time `json:"created_at"`
	UpdatedAT     time.Time `json:"updated_at"`
}

func (wa WhitelistedAddresses) Blockchain() Blockchains {
	var blockchain Blockchains
	config.DataBase.Where("key = ?", wa.BlockchainKey).First(&blockchain)

	return blockchain
}
