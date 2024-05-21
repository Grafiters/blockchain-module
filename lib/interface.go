package lib

type HeightParams struct {
	Height int64 `json:"height"`
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
	Amount          string `json:"amount,omitempty"`
	ContractAddress string `json:"contractAddress,omitempty"`
	Type            string `json:"type,omitempty"`
}

type TranscationConvertion struct {
	Hash          string `json:"hash,omitempty"`
	Amount        string `json:"amount,omitempty"`
	FromAddresses string `json:"from_addresses,omitempty"`
	ToAddress     string `json:"to_address,omitempty"`
	Txout         int64  `json:"txout,omitempty"`
	BlockNumber   int64  `json:"block_number,omitempty"`
	CurrencyID    string `json:"currency_id,omitempty"`
	Status        string `json:"status,omitempty"`
}
