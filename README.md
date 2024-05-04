# Trust Wallet Ethereum Blockchain Parser

The Trust Wallet Ethereum Blockchain Parser is a tool designed to parse Ethereum blockchain data and provide real-time updates on transactions for subscribed addresses. 
It consists of a **parser** component responsible for interacting with the Ethereum JSON-RPC interface and a **processor** component that continuously fetches and updates blockchain data asynchronously.


## Features

- **Real-Time Updates**: Receive real-time updates on inbound and outbound transactions for subscribed Ethereum addresses.
- **Easy Integration**: Integrate Trust Wallet's Ethereum blockchain parser into your application with ease using the provided API endpoints.
- **Customizable**: Extend the functionality of Trust Wallet by customizing the parser and processor components to suit your specific requirements.


## Components

### 1. Parser

The `parser` package provides an interface for interacting with the Ethereum blockchain. It includes methods for fetching the current block number, subscribing to addresses, and retrieving transactions.

### 2. Processor

The `processor` package is responsible for asynchronously processing Ethereum blockchain data. It continuously fetches the latest block number, updates the parser with new data, and sends notifications for transactions related to subscribed addresses.


## Installation

To use the Trust Wallet Ethereum Blockchain Parser in your project, follow these steps:

1. Install Go (if not already installed) by following the instructions [here](https://golang.org/doc/install).

2. Clone the Trust Wallet repository to your local machine: git clone https://github.com/dr-zeus13/trust-wallet.git

3. Navigate to the `trust-wallet` directory: cd trust-wallet

4. Build and run the main package: go run main.go


## Usage

Once the Trust Wallet Ethereum Blockchain Parser is running, you can interact with it using the provided API endpoints:

- **Subscribe**: Use the `/subscribe` endpoint to subscribe to Ethereum addresses and receive real-time transaction updates.

POST /subscribe
Body: {"address": "0x1234567890..."}

curl --location --request GET 'localhost:8081/subscribe' \
--header 'Content-Type: application/json' \
--data '{
    "address": "0x1234567890123456789012345678901234567890"
}'

## Please note:

Appropriate comments have been added, some keypoints can be seached in the repository by `NOTE` or `TODO`.
HTTP server is running on port **8081**
