package parser

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// EthereumRPCClient represents a client for interacting with the Ethereum JSON-RPC endpoint.
type EthereumRPCClient struct {
	URL string
	// multiple params regarding client can be appended viz.: SSL, TLS, Certificates, etc.
}

// Request sends a JSON-RPC request to the Ethereum node.
func (client *EthereumRPCClient) Request(
	method string, params []interface{},
) (map[string]interface{}, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(client.URL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
