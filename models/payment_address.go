package models

import "github.com/nusa-exchange/finex/config"

type PaymentAddress struct {
	ID               string `json:"id"`
	Address          string `json:"address"`
	SecretEncrypted  string `json:"secret_encrypted"`
	DetailsEncrypted string `json:"details_encrypted"`
	MemberID         string `json:"member_id"`
	WalletID         string `json:"wallet_id"`
	Remote           string `json:"remote"`
	BlockchainKey    string `json:"blockchain_key"`
	CreatedAT        string `json:"created_at"`
	UpdatedAT        string `json:"updated_at"`
}

func (pa *PaymentAddress) Secret() string {
	secret, err := config.Vault.DecryptValue("_payment_addresses_secret", pa.SecretEncrypted)
	if err != nil {
		config.Logger.Error(err)
	}

	return secret
}

func (pa *PaymentAddress) Details() string {
	details, err := config.Vault.DecryptValue("_payment_addresses_details", pa.DetailsEncrypted)
	if err != nil {
		config.Logger.Error(err)
	}

	return details
}

func (pa *PaymentAddress) Member() *Member {
	var member *Member
	config.DataBase.Where("id = ?", pa.MemberID).First(&member)

	return member
}

func (pa *PaymentAddress) Blockchain() *Blockchains {
	var blockchain *Blockchains
	config.DataBase.Where("key = ?", pa.BlockchainKey).First(&blockchain)

	return blockchain
}

func (pa *PaymentAddress) Wallet() *Wallet {
	var wallet *Wallet
	config.DataBase.Where("id = ?", pa.WalletID).First(&wallet)

	return wallet
}
