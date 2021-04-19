# Bitcoin Balance Api

An API made in golang that returns the total confirmed and unconfirmed balances to a bitcoin address based on UTXO list.

## Methodology

This API consumes blockbook's(https://blockbook-bitcoin.tronwallet.me/api/v2/utxo/{btc-address}) API to check each bitcoin transaction's confirmations. Therefore, if transaction's confirmations is less than 2, the balance will be unconfirmed, otherwise will be the opposite. Finally, the API will return the total confirmed and unconfirmed balances.

## How to run

Clone the project: 
`git clone https://github.com/Garubieru/btc-balance-api.git`

Build the application:
`go build src/main.go`

Execute:
`./main`

At this time, you have a RESTful API server running at [http://localhost:8000](http://localhost:8000). It provides the following endpoints:

- `GET /balance/{address}`: Returns the total confirmed and unconfirmed balances of a bitcoin address.
