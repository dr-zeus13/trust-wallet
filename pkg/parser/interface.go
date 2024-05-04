package parser

import "github.com/trust-wallet/ethrpc/pkg/models"

// Parser represents the Ethereum blockchain parser interface.
type Parser interface {
	// last parsed block
	GetCurrentBlock() int

	// add address to observer (websocket subscription can be done)
	Subscribe(address string) bool

	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []models.Transaction
}
