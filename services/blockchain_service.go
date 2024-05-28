package services

import (
	"github.com/nusa-exchange/finex/config"
	"github.com/nusa-exchange/finex/lib"
	"github.com/nusa-exchange/finex/models"
)

type BlockchainService struct {
	Blockchain           *models.Blockchains
	BlockchainCurrencies []*models.BlockchainCurrencies
	Currencies           []string
	WhitelistedAddresses []*models.WhitelistedAddresses
	Adapter              lib.LibBlockchain
}

func NewBlockchainService(blockchain *models.Blockchains) BlockchainService {
	var blockchain_currencies []*models.BlockchainCurrencies
	var currencies []*models.Currency
	currency := make([]string, 0)

	blockchain_api_setting := make([]models.ToBlockchainAPISetting, 0)

	config.DataBase.Where("blockchain_key = ?", blockchain.Key).Find(&blockchain_currencies)
	config.DataBase.Where("blockchain_key = ?", blockchain.Key).Find(&currencies)

	for _, bc := range blockchain_currencies {
		blockchain_api_setting = append(blockchain_api_setting, bc.ToAPISetting())
		if !isInArray(currency, bc.CurrencyID) {
			currency = append(currency, bc.CurrencyID)
		}
	}

	adapter := ClientServer(blockchain, blockchain_api_setting, blockchain.WhitelistedAddresses())

	return BlockchainService{
		Blockchain:           blockchain,
		BlockchainCurrencies: blockchain_currencies,
		Currencies:           currency,
		WhitelistedAddresses: blockchain.WhitelistedAddresses(),
		Adapter:              adapter,
	}
}

func (bc_service *BlockchainService) Height() int {
	return 16
}

func (bc_service *BlockchainService) ProcessBlock(height int64) error {
	_, err := bc_service.Adapter.FetchBlocks(height)
	if err != nil {
		config.Logger.Error(err)
		return err
	}

	return nil
}

func (bc_service *BlockchainService) Deposit(data []*lib.TranscationConvertion) error {
	var wallet []*models.Wallet
	config.DataBase.Where("kind = ?", models.DepositKindWallet).First(&wallet)

	return nil
}
