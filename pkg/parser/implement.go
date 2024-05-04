package parser

import "github.com/trust-wallet/ethrpc/pkg/models"

// NOTE: Making structs, functions and params `PUBLIC` for simplicity

// EthereumParser implements Parser
var _ Parser = &EthereumParser{}

// EthereumParser represents the implementation of the Parser interface.
type EthereumParser struct {
	Client          *EthereumRPCClient
	currentBlock    int
	subscribedAddrs map[string]bool                 // List of subscribed addresses
	transactions    map[string][]models.Transaction // Mapping of address to its transactions
}

// NewEthereumParser creates a new instance of EthereumParser.
func NewEthereumParser(url string) *EthereumParser {
	return &EthereumParser{
		Client:          &EthereumRPCClient{URL: url},
		currentBlock:    0,
		subscribedAddrs: make(map[string]bool),
		transactions:    make(map[string][]models.Transaction),
	}
}

// GetCurrentBlock returns the last parsed block.
func (ep *EthereumParser) GetCurrentBlock() int {
	return ep.currentBlock
}

// SetCurrentBlock sets the last parsed block.
func (ep *EthereumParser) SetCurrentBlock(currentBlock int) {
	ep.currentBlock = currentBlock
}

// Subscribe adds an address to the observer.
func (ep *EthereumParser) Subscribe(address string) bool {
	ep.subscribedAddrs[address] = true
	return true // Assuming always successful for simplicity
}

// GetSubscribedAddresses returns map of addresses subscribed
func (ep *EthereumParser) GetSubscribedAddresses() map[string]bool {
	return ep.subscribedAddrs
}

// GetTransactions returns a list of inbound or outbound transactions for an address.
func (ep *EthereumParser) GetTransactions(address string) []models.Transaction {
	return ep.transactions[address]
}
