package services

import (
	"github.com/nusa-exchange/finex/config"
	"github.com/nusa-exchange/finex/models"
)

type BlockchainService struct {
	Blockchain           *models.Blockchains
	BlockchainCurrencies []*models.BlockchainCurrencies
	Currencies           []string
	WhitelistedAddresses []*models.WhitelistedAddresses
}

func NewBlockchainService(blockchain *models.Blockchains) BlockchainService {
	var blockchain_currencies []*models.BlockchainCurrencies
	var currencies []string

	config.DataBase.Where("blockchain_key = ?", blockchain.Key).Find(&blockchain_currencies)
	config.DataBase.Where("blockchain_key = ?", blockchain.Key).Scan(&currencies)

	return BlockchainService{
		Blockchain:           blockchain,
		BlockchainCurrencies: blockchain_currencies,
		Currencies:           currencies,
		WhitelistedAddresses: blockchain.WhitelistedAddresses(),
	}
}

func (bc_service *BlockchainService) Height() int {
	return 16
}

func (bc_service *BlockchainService) ProcessBlock(height int64) {

}
