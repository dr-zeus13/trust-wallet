package models

// Transaction represents a transaction on the Ethereum blockchain.
// Note: Assuming it is a simple `Send` transaction else params: From, To, Value
// can be extracted to child struct of Transaction
type Transaction struct {
	TxHash      string `json:"tx_hash"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	BlockHash   string `json:"block_hash"`
	BlockNumber string `json:"block_number"`
	Transaction string `json:"transaction_index"`
}
