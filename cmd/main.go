package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/trust-wallet/ethrpc/pkg/parser"
	"github.com/trust-wallet/ethrpc/pkg/processor"
)

func main() {
	ethereumParser := parser.NewEthereumParser("https://cloudflare-eth.com")

	processor := processor.NewProcessor(ethereumParser, 10*time.Second) // Update every 10 seconds
	processor.Start()
	defer processor.Stop()
	fmt.Println("Processor started...")

	// Basis the user's subscription request,
	// class: `Processor` will send notifications to user asynchronously
	http.Handle("/subscribe", http.HandlerFunc(subscribeHandler(ethereumParser)))

	fmt.Println("Server started listening at port: 8081.....")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println(err)
	}
}

// subscribeHandler: can be used in an API, to subscribe events via websocket
func subscribeHandler(ep *parser.EthereumParser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			Address string `json:"address"`
		}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if data.Address == "" {
			http.Error(w, "Address not provided", http.StatusBadRequest)
			return
		}
		ep.Subscribe(data.Address)
		fmt.Fprintf(w, "Subscribed to address: %s", data.Address)
	}
}
