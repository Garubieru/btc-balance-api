package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Transactions []struct {
	Txid          string `json:"txid"`
	Vout          int    `json:"vout"`
	Value         string `json:"value"`
	Confirmations int    `json:"confirmations"`
	Height        int    `json:"height,omitempty"`
}

type TotalBalance struct {
	Confirmed   int `json:"confirmed"`
	Unconfirmed int `json:"unconfirmed"`
}

func getUtxoList(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatalln(err2)
	}

	return body
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	url := fmt.Sprintf("https://blockbook-bitcoin.tronwallet.me/api/v2/utxo/%s", params["address"])
	body := getUtxoList(url)

	var transactions = new(Transactions)

	var totalBalance TotalBalance

	err := json.Unmarshal(body, &transactions)
	if err != nil {
		json.NewEncoder(w).Encode(totalBalance)
		return
	}

	for _, transaction := range *transactions {
		if transaction.Confirmations < 2 {
			totalBalance.Unconfirmed += 1
		} else {
			totalBalance.Confirmed += 1
		}
	}
	json.NewEncoder(w).Encode(totalBalance)
}

func main() {
	router := mux.NewRouter()
	port := ":8000"

	router.HandleFunc("/balance/{address}", getBalance).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Println("Listening on port ", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
