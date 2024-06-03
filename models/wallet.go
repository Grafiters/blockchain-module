package models

import "github.com/nusa-exchange/finex/config"

type KindWallet int16

var (
	HotKindWallet     KindWallet = 310
	FeeKindWallet     KindWallet = 200
	DepositKindWallet KindWallet = 100
)

type Wallet struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	Address           string     `json:"address"`
	Status            string     `json:"status"`
	CreatedAT         string     `json:"created_at"`
	UpdatedAT         string     `json:"updated_at"`
	Gateway           string     `json:"gateway"`
	MaxBalance        string     `json:"max_balance"`
	BlockchainKey     string     `json:"blockchain_key"`
	Kind              KindWallet `json:"kind"`
	SettingsEncrypted string     `json:"settings_encrypted"`
	Balance           string     `json:"balance"`
	PlainSettings     string     `json:"plain_settings"`
}

func (w *Wallet) Blockchain() *Blockchains {
	var blockchain *Blockchains
	config.DataBase.Where("key = ?", w.BlockchainKey)

	return blockchain
}

func (w *Wallet) Settings() string {
	details, err := config.Vault.DecryptValue("_wallets_settings", w.SettingsEncrypted)
	if err != nil {
		config.Logger.Error(err)
	}

	return details
}

func (w *Wallet) CurrencyWalelt(currency string) *CurrencyWallet {
	var cw *CurrencyWallet
	config.DataBase.Where("currency_id = ?", currency).First(&cw)
	return cw
}
