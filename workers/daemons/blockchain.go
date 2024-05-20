package daemons

import (
	"time"

	"github.com/nusa-exchange/finex/config"
	"github.com/nusa-exchange/finex/models"
	"github.com/nusa-exchange/finex/services"
)

type Blockchain struct {
	Running     bool
	Blockchains []*models.Blockchains
}

func NewBlockchain() *Blockchain {
	var blockchains []*models.Blockchains
	config.DataBase.Where("status = ?", "active").Find(&blockchains)

	return &Blockchain{Running: true, Blockchains: blockchains}
}

func (b *Blockchain) Stop() {
	b.Running = false
}

func (b *Blockchain) Start() {
	for b.Running {

		for _, blockchain := range b.Blockchains {
			bc_service := services.NewBlockchainService(blockchain)
			height := bc_service.Height()
			config.Logger.Info(height)
		}

		time.Sleep(1 * time.Second)
	}
}

func (b *Blockchain) Process() {
	for {
		var blockchains []*models.Blockchains
		config.DataBase.Where("status = ?", "active").Find(&blockchains)

		if !b.Running {
			break
		}

		for _, v := range blockchains {
			bc_service := services.NewBlockchainService(v)
			height := bc_service.Height()
			config.Logger.Info(height)
		}
	}
}
