package processor

import (
	"fmt"
	"strconv"
	"time"

	"github.com/trust-wallet/ethrpc/pkg/parser"
)

// NOTE: this file is responsible for running asynchronously
// querying blockchain for its latest height
// updating: currentBlockNumber and transactions list basis user addresses

// Processor represents the component responsible for processing Ethereum blocks.
type Processor struct {
	parser  *parser.EthereumParser
	ticker  *time.Ticker
	running bool
}

// NewProcessor creates a new instance of Processor.
func NewProcessor(parser *parser.EthereumParser, interval time.Duration) *Processor {
	return &Processor{
		parser: parser,
		ticker: time.NewTicker(interval),
	}
}

// Start starts the processor.
func (p *Processor) Start() {
	if p.running {
		return
	}
	p.running = true
	go p.processBlocks()
}

// Stop stops the processor.
func (p *Processor) Stop() {
	if !p.running {
		return
	}
	p.running = false
	p.ticker.Stop()
}

func (p *Processor) processBlocks() {
	for range p.ticker.C {
		// Fetch the current block number
		result, err := p.parser.Client.Request("eth_blockNumber", nil)
		if err != nil {
			fmt.Println("Error fetching current block number:", err)
			continue
		}
		blockNumber := result["result"].(string)

		block, _ := strconv.Atoi(blockNumber)
		p.parser.SetCurrentBlock(block)

		// TODO: Similarly parse the transactions of each block
		// and add them to the mermory store

		// Process transactions for subscribed addresses
		// to be sent to user via Notification Service
		for address := range p.parser.GetSubscribedAddresses() {
			p.processTransactionsForAddress(address)
		}
	}
}

func (p *Processor) processTransactionsForAddress(address string) {
	// TODO: Implement the logic to fetch transactions for the given address
	// and update the parser's transactions mapping.
	// This function will depend on your specific requirements and the
	// JSON-RPC methods available for fetching transactions.
	fmt.Println("Processing transactions for address: ", address)

	// Send to user via Notification Service
	transactions := p.parser.GetTransactions(address)
	fmt.Printf("\nList of transactions to be sent to user via websocket for address %s are: %v",
		address, transactions)
}
