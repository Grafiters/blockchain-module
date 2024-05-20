package models

import (
	"time"

	"github.com/nusa-exchange/finex/config"
	"github.com/shopspring/decimal"
)

type Blockchains struct {
	ID                  string          `json:"id" gorm:"primaryKey"`
	Key                 string          `json:"key"`
	Name                string          `json:"name"`
	Client              string          `json:"client"`
	Height              int64           `json:"height"`
	ExplorerAddress     string          `json:"explorer_address"`
	ExplorerTransaction string          `json:"explorer_transaction"`
	MinConfirmations    int64           `json:"min_confirmations"`
	Status              string          `json:"status"`
	Description         string          `json:"description"`
	Warning             string          `json:"warning"`
	Protocol            string          `json:"protocol"`
	CollectionGasSpeed  string          `json:"collection_gas_speed"`
	WithdrawalGasSpeed  string          `json:"withdrawal_gas_speed"`
	ServerEncrypted     string          `json:"server_encrypted"`
	CreatedAT           time.Time       `json:"created_at"`
	UpdatedAT           time.Time       `json:"updated_at"`
	MinDepositAmount    decimal.Decimal `json:"min_deposit_amount"`
	WithdrawFee         decimal.Decimal `json:"withdraw_fee"`
	MinWithdrawAmount   decimal.Decimal `json:"min_withdraw_amount"`
}

func (b Blockchains) WhitelistedAddresses() []*WhitelistedAddresses {
	var whitelist []*WhitelistedAddresses
	config.DataBase.Where("blockchain_key = ?", b.Key).Find(&whitelist)

	return whitelist
}
