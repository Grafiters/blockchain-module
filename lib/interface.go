package lib

import "github.com/shopspring/decimal"

type HeightParams struct {
	Height int64 `json:"height"`
}

type ReceiptParams struct {
	Hash string `json:"hash"`
}

type TransactionBlockchains struct {
	Transactions []*TransactionsBlockchain `json:"transactions,omitempty"`
}

type TransactionsBlockchain struct {
	Txid            string `json:"txid,omitempty"`
	From            string `json:"from,omitempty"`
	GasPrice        string `json:"gasPrice,omitempty"`
	GasLimit        string `json:"gasLimit,omitempty"`
	To              string `json:"to,omitempty"`
	Amount          int64  `json:"amount,omitempty"`
	ContractAddress string `json:"contractAddress,omitempty"`
	Type            string `json:"type,omitempty"`
}

type TranscationConvertion struct {
	Hash          string          `json:"hash,omitempty"`
	Amount        decimal.Decimal `json:"amount,omitempty"`
	FromAddresses string          `json:"from_addresses,omitempty"`
	ToAddress     string          `json:"to_address,omitempty"`
	Txout         int64           `json:"txout,omitempty"`
	BlockNumber   int64           `json:"block_number,omitempty"`
	CurrencyID    string          `json:"currency_id,omitempty"`
	Status        string          `json:"status,omitempty"`
}

type TransactionReceipt struct {
	Type              string        `json:"type"`
	From              string        `json:"from"`
	To                string        `json:"to"`
	Status            bool          `json:"status"`
	CumulaticeGasUsed int64         `json:"cumulatice_gas_used"`
	LogsBloom         string        `json:"logs_bloom"`
	Logs              []interface{} `json:"logs"`
	TransactionHash   string        `json:"transaction_hash"`
	ContractAddress   string        `json:"contract_address"`
	GasUsed           int64         `json:"gas_used"`
	BlockHash         string        `json:"block_hash"`
	BlockNumber       int64         `json:"block_number"`
	TransactionIndex  int32         `json:"transaction_index"`
	EffectiveGasPrice int64         `json:"effective_gas_price"`
}
