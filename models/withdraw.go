package models

import (
	"time"

	"github.com/nusa-exchange/finex/config"
	"github.com/shopspring/decimal"
)

type WithdrawAasmStateType string

var (
	WithdrawAcceptedState      WithdrawAasmStateType = "accepted"
	WithdrawProcessingState    WithdrawAasmStateType = "processing"
	WithdrawFeeProcessingState WithdrawAasmStateType = "fee_processing"
	WithdrawSkippedState       WithdrawAasmStateType = "skipped"
	WithdrawCanceledState      WithdrawAasmStateType = "canceled"
	WithdrawRejectedState      WithdrawAasmStateType = "rejected"
	WithdrawConfirmingState    WithdrawAasmStateType = "confirming"
	WithdrawSucceedState       WithdrawAasmStateType = "succeed"
	WithdrawErroredState       WithdrawAasmStateType = "errored"
)

type Withdraw struct {
	ID            int64                 `json:"id"`
	MemberID      int64                 `json:"member_id"`
	CurrencyID    string                `json:"currency_id"`
	Amount        decimal.Decimal       `json:"amount"`
	Fee           decimal.Decimal       `json:"fee" gorm:"default:0.0"`
	CreatedAT     time.Time             `json:"created_at"`
	UpdatedAT     time.Time             `json:"updated_at"`
	CompletedAT   time.Time             `json:"completed_at" gorm:"default:null"`
	Txid          string                `json:"txid"`
	AasmState     WithdrawAasmStateType `json:"aasm_state" gorm:"default:accepted"`
	Sum           decimal.Decimal       `json:"sum"`
	Type          string                `json:"type"`
	Tid           string                `json:"tid"`
	Rid           string                `json:"rid"`
	BlockNumber   int64                 `json:"block_number"`
	Note          string                `json:"note"`
	Error         string                `json:"error"`
	BeneficiaryID int64                 `json:"beneficiary_id"`
	TransferType  string                `json:"transfer_type"`
	Metadata      string                `json:"metadata"`
	RemoteID      string                `json:"remote_id"`
	BlockchainKey string                `json:"blockchain_key"`
}

func (d *Withdraw) Members() *Member {
	var member *Member
	config.DataBase.Where("id = ?", d.MemberID).First(&member)

	return member
}

func (d *Withdraw) Currency() *Currency {
	var currency *Currency
	config.DataBase.Where("id = ?", d.CurrencyID).First(&currency)

	return currency
}

func (d *Withdraw) Blockchain() *Blockchains {
	var blockchain *Blockchains
	config.DataBase.Where("key = ?", d.BlockchainKey).First(&blockchain)

	return blockchain
}

func (d *Withdraw) Beneficiary() *Beneficiary {
	var blockchain *Beneficiary
	config.DataBase.Where("key = ?", d.BlockchainKey).First(&blockchain)

	return blockchain
}
