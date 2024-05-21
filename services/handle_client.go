package services

import (
	"github.com/nusa-exchange/finex/lib"
	"github.com/nusa-exchange/finex/lib/ether"
	"github.com/nusa-exchange/finex/models"
)

func ClientServer(blockchain *models.Blockchains, currencies []models.ToBlockchainAPISetting, whitelist []*models.WhitelistedAddresses) lib.LibBlockchain {
	switch blockchain.Client {
	case "ether":
		return ether.NewBlockchain(blockchain.Server(), currencies, whitelist)
	default:
		return nil
	}
}
