package ether

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nusa-exchange/finex/config"
	"github.com/nusa-exchange/finex/controllers/helpers"
	"github.com/nusa-exchange/finex/lib"
	"github.com/nusa-exchange/finex/models"
	"github.com/shopspring/decimal"
)

type EtherBlockchain struct{}

type BlockchainLibrary struct {
	Server               string
	Coin                 models.ToBlockchainAPISetting
	SmartContract        []models.ToBlockchainAPISetting
	WhitelistedAddresses []*models.WhitelistedAddresses
	Client               lib.ClientLib
}

func NewBlockchain(server string, currencies []models.ToBlockchainAPISetting, whitelist []*models.WhitelistedAddresses) *BlockchainLibrary {
	var coin models.ToBlockchainAPISetting
	contract := make([]models.ToBlockchainAPISetting, 0)

	for _, c := range currencies {
		if len(c.Options.ContractAddress) > 0 {
			contract = append(contract, c)
		} else {
			coin = c
		}
	}

	client := lib.NewClient(server)

	return &BlockchainLibrary{
		Server:               server,
		Coin:                 coin,
		SmartContract:        contract,
		WhitelistedAddresses: whitelist,
		Client:               client,
	}
}

func (bl *BlockchainLibrary) FetchBlocks(height int64) ([]*lib.TranscationConvertion, error) {
	var transaction *lib.TransactionBlockchains
	transactions := make([]*lib.TranscationConvertion, 0)

	body := lib.HeightParams{
		Height: height,
	}

	requestBody, err := json.Marshal(body)
	if err != nil {
		config.Logger.Error("Failed to marshal object :", err)
		return nil, err
	}

	data, err := bl.Client.Post("fetch-block", *bytes.NewBuffer(requestBody))
	if err != nil {
		config.Logger.Error(err)
		return nil, err
	}

	recordJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling data to JSON:", err)
		return nil, err
	}

	err = json.Unmarshal(recordJSON, &transaction)
	if err != nil {
		return nil, err
	}

	for _, tx := range transaction.Transactions {
		if tx.Type == "Ether" {
			tx_data := &lib.TranscationConvertion{
				Hash:          tx.Txid,
				Amount:        helpers.ConvertFromBase(decimal.NewFromInt(tx.Amount), bl.Coin.BaseFactor),
				FromAddresses: tx.From,
				ToAddress:     tx.To,
				Txout:         1,
				BlockNumber:   height,
				CurrencyID:    bl.Coin.ID,
				Status:        "pending",
			}

			transactions = append(transactions, tx_data)
		} else {
			if len(bl.SmartContract) > 0 {
				for _, sc := range bl.SmartContract {
					if strings.ToLower(tx.ContractAddress) == strings.ToLower(sc.Options.ContractAddress) {
						tx_data := &lib.TranscationConvertion{
							Hash:          tx.Txid,
							Amount:        helpers.ConvertFromBase(decimal.NewFromInt(tx.Amount), sc.BaseFactor),
							FromAddresses: tx.From,
							ToAddress:     tx.To,
							Txout:         1,
							BlockNumber:   height,
							CurrencyID:    sc.ID,
							Status:        "pending",
						}

						transactions = append(transactions, tx_data)
					}
				}
			}
		}
	}

	return transactions, nil
}

func (bl *BlockchainLibrary) FetchTransaction(hash string) (*lib.TransactionReceipt, error) {
	var receipt *lib.TransactionReceipt
	body := lib.ReceiptParams{
		Hash: hash,
	}

	requestBody, err := json.Marshal(body)
	if err != nil {
		config.Logger.Error("Failed to marshal object :", err)
		return nil, err
	}

	data, err := bl.Client.Post("get-transaction", *bytes.NewBuffer(requestBody))
	if err != nil {
		config.Logger.Error(err)
		return nil, err
	}

	recordJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling data to JSON:", err)
		return nil, err
	}

	err = json.Unmarshal(recordJSON, &receipt)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}
