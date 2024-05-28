package models

import "github.com/nusa-exchange/finex/config"

type CurrencyWallet struct {
	CurrencyID string `json:"currency_id"`
	WalletID   int64  `json:"wallet_id"`
}

func (cw *CurrencyWallet) Currency() *Currency {
	var currency *Currency
	config.DataBase.Where("id = ?", cw.CurrencyID).First(&currency)

	return currency
}

func (cw *CurrencyWallet) Wallet() *Wallet {
	var wallet *Wallet
	config.DataBase.Where("id = ?", cw.WalletID).First(&wallet)

	return wallet
}
