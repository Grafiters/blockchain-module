package services

import (
	"github.com/nusa-exchange/finex/config"
	"github.com/nusa-exchange/finex/lib"
	"github.com/nusa-exchange/finex/models"
	"github.com/shopspring/decimal"
)

type BlockchainService struct {
	Blockchain           *models.Blockchains
	BlockchainCurrencies []*models.BlockchainCurrencies
	Currencies           []string
	WhitelistedAddresses []*models.WhitelistedAddresses
	Adapter              lib.LibBlockchain
}

type DepositWithdraw struct {
	Deposit  []*lib.TranscationConvertion
	Withdraw []*lib.TranscationConvertion
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
	data, err := bc_service.Adapter.FetchBlocks(height)
	if err != nil {
		config.Logger.Error(err)
		return err
	}

	filteringTx := bc_service.FilteringData(data)
	config.Logger.Info(filteringTx)

	return nil
}

func (bc_service *BlockchainService) FilteringData(data []*lib.TranscationConvertion) *DepositWithdraw {
	deposit_data := make([]*lib.TranscationConvertion, 0)
	withdraw_data := make([]*lib.TranscationConvertion, 0)
	var wallet *models.Wallet

	for _, i := range data {
		var currency_wallet *models.CurrencyWallet
		config.DataBase.Where("blockchain_key = ? AND kind = ?", bc_service.Blockchain.Key, models.DepositKindWallet).First(&wallet)
		currency_wallet = wallet.CurrencyWalelt(i.CurrencyID)

		var address_depo *models.PaymentAddress
		config.DataBase.Where("address = ? AND blockchain_key = ? AND wallet_id = ?", i.ToAddress, bc_service.Blockchain.Key, currency_wallet.WalletID).First(&address_depo)

		if address_depo != nil {
			deposit_data = append(deposit_data, i)
		}

		config.DataBase.Where("blockchain_key = ? AND kind = ?", bc_service.Blockchain.Key, models.HotKindWallet).First(&wallet)
		currency_wallet = wallet.CurrencyWalelt(i.CurrencyID)
		var address_withdraw *models.PaymentAddress
		config.DataBase.Where("address = ? AND blockchain_key = ? AND wallet_id = ?", i.FromAddresses, bc_service.Blockchain.Key, currency_wallet.WalletID).First(&address_withdraw)

		if address_withdraw != nil {
			withdraw_data = append(withdraw_data, i)
		}
	}

	return &DepositWithdraw{
		Deposit:  deposit_data,
		Withdraw: withdraw_data,
	}
}

func (bc_service *BlockchainService) UpdateDeposit(transaction *lib.TranscationConvertion) error {
	var (
		bc_currency     *models.BlockchainCurrencies
		wallet          *models.Wallet
		currency_wallet *models.CurrencyWallet
		payment_address *models.PaymentAddress
		transactions    *models.Transactions
	)
	config.DataBase.Where("blockchain_key = ? AND currency_id = ?", bc_service.Blockchain.Key, transaction.CurrencyID).First(&bc_currency)

	if transaction.Amount.LessThan(bc_currency.MinDepositAmount) {
		config.Logger.Info("Skip deposit with hash ", transaction.Hash, " with amount ", transaction.Amount, " to ", transaction.ToAddress, " in block number ", transaction.BlockNumber, " cause aomunt is less then min deposit")
		return nil
	}

	receipt, err := bc_service.Adapter.FetchTransaction(transaction.Hash)
	if err != nil || !receipt.Status {
		config.Logger.Info("Skip update Deposit hash", transaction.Hash, "not found or network is still : pending")
		return err
	}

	config.DataBase.Where("blockchain_key = ? AND kind = ?", bc_service.Blockchain.Key, models.DepositKindWallet).First(&wallet)
	currency_wallet = wallet.CurrencyWalelt(transaction.CurrencyID)

	config.DataBase.Where("wallet_id = ? AND address = ?", currency_wallet.WalletID, transaction.ToAddress).First(&payment_address)
	if wallet == nil {
		config.Logger.Info("Skip deposit with hash ", transaction.Hash, "cause address with wallet deposit is disabled")
		return nil
	}

	config.DataBase.Where("txid = ? AND referance_type = ?", transaction.Hash, models.ReferenceDeposit).First(&transactions)
	if transactions != nil {
		config.Logger.Info("Skip deposit with hash ", transaction.Hash, " transaction already saved on database")
		return nil
	}

	if transaction.FromAddresses == "" {
		config.Logger.Info("Skip deposit with hash ", transaction.Hash, " transaction source from address is null")
		return nil
	}

	return nil
}

func (bc_service *BlockchainService) UpdateWithdraw(transaction *lib.TranscationConvertion) error {
	var withdraw *models.Withdraw
	var tx *models.Transactions
	status := models.TransactinPending
	config.DataBase.Where("hash = ? AND blockcahain_key = ? AND currency_id = ?", transaction.Hash, bc_service.Blockchain.Key, transaction.CurrencyID).First(&withdraw)

	if withdraw == nil {
		config.Logger.Info("Skip update withdrawal: ", transaction.Hash)
		return nil
	}

	config.DataBase.Model(&withdraw).Update("block_number", transaction.BlockNumber)

	receipt, err := bc_service.Adapter.FetchTransaction(transaction.Hash)
	if err != nil || !receipt.Status {
		config.Logger.Info("Skip update withdrawal hash", transaction.Hash, "not found or network is still : pending")
		return err
	}

	config.DataBase.Where("txid = ?", transaction.Hash).First(&tx)
	if tx == nil {
		config.Logger.Info("Data ", transaction.Hash, "is not registered on dataabse")
		return nil
	}

	if receipt.Status {
		status = models.TransactinSuccess
	}

	if !receipt.Status {
		config.Logger.Info("Skip withdraw ", transaction.Hash, "cause from network")
		return nil
	}

	config.DataBase.Model(&tx).Updates(&models.Transactions{
		Fee:    decimal.NewFromInt(receipt.EffectiveGasPrice),
		Status: status,
	})

	return nil
}
