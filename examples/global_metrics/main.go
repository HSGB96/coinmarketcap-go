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

	metrics, err := client.GlobalMetrics().QuotesLatest(map[string]string{
		"convert": "USD",
	})

	if err != nil {
		log.Fatal(err)
	}

	data := metrics["data"].(map[string]interface{})
	quote := data["quote"].(map[string]interface{})
	usd := quote["USD"].(map[string]interface{})

	fmt.Println("Global Cryptocurrency Market Metrics")
	fmt.Println("====================================")
	fmt.Printf("Total Market Cap:        $%15.0f\n", usd["total_market_cap"].(float64))
	fmt.Printf("Total Volume 24h:        $%15.0f\n", usd["total_volume_24h"].(float64))
	fmt.Printf("BTC Dominance:           %15.2f%%\n", data["btc_dominance"].(float64))
	fmt.Printf("ETH Dominance:           %15.2f%%\n", data["eth_dominance"].(float64))
	fmt.Printf("Active Cryptocurrencies: %15.0f\n", data["active_cryptocurrencies"].(float64))
	fmt.Printf("Active Exchanges:        %15.0f\n", data["active_exchanges"].(float64))
	fmt.Printf("Active Market Pairs:     %15.0f\n", data["active_market_pairs"].(float64))
}
