package models

import (
	"encoding/json"
	"time"

	"github.com/nusa-exchange/finex/config"
	"github.com/shopspring/decimal"
)

type BlockchainCurrencies struct {
	ID                    int64           `json:"id"`
	CurrencyID            string          `json:"currency_id"`
	BlockchainKey         string          `json:"blockchain_key"`
	ParentID              string          `json:"parent_id"`
	DepositFee            decimal.Decimal `json:"deposit_fee"`
	MinDepositAmount      decimal.Decimal `json:"min_deposit_amount"`
	MinCollectionAmount   decimal.Decimal `json:"min_collection_amount"`
	WithdrawFee           decimal.Decimal `json:"withdraw_fee"`
	MinWithdrawAmount     decimal.Decimal `json:"min_withdraw_amount"`
	DepositEnabled        bool            `json:"deposit_enabled"`
	WithdrawalEnabled     bool            `json:"withdrawal_enabled"`
	AutoUpdateFeesEnabled bool            `json:"auto_update_fees_enabled"`
	BaseFactor            int64           `json:"base_factor"`
	Status                string          `json:"status"`
	Options               []byte          `json:"options" gorm:"type:json"`
	CreatedAt             time.Time       `json:"created_at"`
	UpdatedAt             time.Time       `json:"updated_at"`
}

type Options struct {
	GasLimit        string `json:"gas_limit,omitempty"`
	GasPrice        string `json:"gas_price,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"`
}

type ToBlockchainAPISetting struct {
	ID                  string          `json:"currency"`
	BaseFactor          int64           `json:"base_factor"`
	MinCollectionAmount decimal.Decimal `json:"min_collection_amount"`
	Options             Options         `json:"options"`
}

func (bc BlockchainCurrencies) Currency() *Currency {
	var currencies *Currency
	config.DataBase.Where("id = ?", bc.CurrencyID).First(&currencies)
	return currencies
}

func (bc BlockchainCurrencies) Blockchain() *Blockchains {
	var blockchain *Blockchains
	config.DataBase.Where("key = ?", bc.BlockchainKey).First(&blockchain)

	return blockchain
}

func (bc *BlockchainCurrencies) ToAPISetting() ToBlockchainAPISetting {
	var optionData *Options
	var apisetting ToBlockchainAPISetting

	err := json.Unmarshal(bc.Options, &optionData)
	if err != nil {
		config.Logger.Error(err)
		return apisetting
	}

	return ToBlockchainAPISetting{
		ID:                  bc.CurrencyID,
		BaseFactor:          bc.BaseFactor,
		MinCollectionAmount: bc.MinCollectionAmount,
		Options:             *optionData,
	}
}
