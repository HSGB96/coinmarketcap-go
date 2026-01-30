package main

import (
	"fmt"
	"log"

	coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
)

func main() {
	client, err := coinmarketcap.NewClientBuilder().
		SetAPIKey("your-api-key-here").
		Build()

	if err != nil {
		log.Fatal(err)
	}

	quotes, err := client.Cryptocurrency().QuotesLatest(map[string]string{
		"symbol":  "BTC,ETH",
		"convert": "USD",
	})

	if err != nil {
		log.Fatal(err)
	}

	data := quotes["data"].(map[string]interface{})

	btc := data["BTC"].(map[string]interface{})
	btcQuote := btc["quote"].(map[string]interface{})
	btcUSD := btcQuote["USD"].(map[string]interface{})

	fmt.Printf("Bitcoin Price: $%.2f\n", btcUSD["price"].(float64))

	eth := data["ETH"].(map[string]interface{})
	ethQuote := eth["quote"].(map[string]interface{})
	ethUSD := ethQuote["USD"].(map[string]interface{})

	fmt.Printf("Ethereum Price: $%.2f\n", ethUSD["price"].(float64))
}
