package models

import (
	"time"

	"github.com/nusa-exchange/finex/config"
	"github.com/nusa-exchange/finex/controllers/entities"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransferType int64
type DepositAasmStateType string

var (
	FiatTransfer   TransferType = 100
	CryptoTransfer TransferType = 200
)

var (
	DepositAcceptedState      DepositAasmStateType = "accepted"
	DepositProcessingState    DepositAasmStateType = "processing"
	DepositFeeProcessingState DepositAasmStateType = "fee_processing"
	DepositSkippedState       DepositAasmStateType = "skipped"
	DepositCanceledState      DepositAasmStateType = "canceled"
	DepositRejectedState      DepositAasmStateType = "rejected"
	DepositSubmittedState     DepositAasmStateType = "submitted"
	DepositCollectedState     DepositAasmStateType = "collected"
	DepositErroredState       DepositAasmStateType = "errored"
)

type Deposit struct {
	ID            int64                `json:"id"`
	MemberID      int64                `json:"member_id"`
	CurrencyID    string               `json:"currency_id"`
	Amount        decimal.Decimal      `json:"amount"`
	Fee           decimal.Decimal      `json:"fee" gorm:"default:0.0"`
	Txid          string               `json:"txid"`
	AasmState     DepositAasmStateType `json:"aasm_state" gorm:"default:accepted"`
	CreatedAT     time.Time            `json:"created_at"`
	UpdatedAT     time.Time            `json:"updated_at"`
	CompletedAT   time.Time            `json:"completed_at" gorm:"default:null"`
	Type          TransferType         `json:"type"`
	Txout         int8                 `json:"txout"`
	Tid           string               `json:"tid"`
	Address       string               `json:"address"`
	BlockNumber   int64                `json:"block_number"`
	Spread        string               `json:"spread" gorm:"default:null"`
	FromAddresses string               `json:"from_addresses"`
	TransferType  int8                 `json:"transfer_type" gorm:"default:null"`
	Error         string               `json:"error" gorm:"null"`
	BlockchainKey string               `json:"blockchain_key"`
}

func (d *Deposit) Accept() error {
	var account *Account
	if d.Currency().Type == TypeCoin {
		err := config.DataBase.Transaction(func(tx *gorm.DB) error {
			if d.AasmState != DepositAcceptedState {
				config.Logger.Info("Deposit already process")
				return nil
			}

			account_tx := tx.Clauses(clause.Locking{Strength: "UPDATE", Table: clause.Table{Name: "accounts"}})
			account_tx.Where("member_id = ? AND currency_id = ?", d.MemberID, d.Currency().ID).FirstOrCreate(&account)
			if err := account.LockFunds(account_tx, d.Amount); err != nil {
				return err
			}

			config.DataBase.Where("id = ?", d.ID).Updates(&Deposit{
				AasmState: DepositProcessingState,
			})
			return nil
		})

		if err != nil {
			config.Logger.Error("can't accept this deposit database may conflict, deposit ID: ", d.ID)
		}

		if err == nil {
			config.RangoClient.EnqueueEvent("private", d.Members().UID, "deposit", d.ToJSON())
		}
	}

	return nil
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

func (d *Deposit) ToJSON() *entities.DepositEntities {
	var confirmation int64
	confirmation = d.Blockchain().Height - d.BlockNumber

	return &entities.DepositEntities{
		ID:            d.ID,
		Currency:      d.CurrencyID,
		BlockchainKey: d.BlockchainKey,
		Protocol:      d.Blockchain().Protocol,
		Warning:       d.Blockchain().Warning,
		Amount:        d.Amount,
		Fee:           d.Fee,
		FromAddresses: d.FromAddresses,
		Txid:          d.Txid,
		Confirmations: confirmation,
		State:         string(d.AasmState),
		TransferType:  string(d.TransferType),
		Error:         d.Error,
		Tid:           d.Tid,
		CreatedAT:     d.CreatedAT,
		CompletedAT:   d.CompletedAT,
	}
}
