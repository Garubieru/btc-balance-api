# Bitcoin Balance Api

An API made in golang that returns the total confirmed and unconfirmed balances to a bitcoin address based on UTXO list.

## Methodology

This API consumes blockbook's(https://blockbook-bitcoin.tronwallet.me/api/v2/utxo/{address}) API to check each transaction confirmations. Therefore, if transaction confirmations is less than 2, the balance will be unconfirmed, otherwise will be the oposite.

## How to run

Install project dependencies:
`go get -v`

Build the application:
`go build src/main.go`

Execute:
`./main`
