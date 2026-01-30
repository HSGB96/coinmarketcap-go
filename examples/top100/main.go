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

	listings, err := client.Cryptocurrency().ListingsLatest(map[string]string{
		"start":    "1",
		"limit":    "100",
		"convert":  "USD",
		"sort":     "market_cap",
		"sort_dir": "desc",
	})

	if err != nil {
		log.Fatal(err)
	}

	data := listings["data"].([]interface{})

	fmt.Println("Top 100 Cryptocurrencies by Market Cap:")
	fmt.Println("========================================")

	for _, item := range data {
		crypto := item.(map[string]interface{})
		quote := crypto["quote"].(map[string]interface{})
		usd := quote["USD"].(map[string]interface{})

		fmt.Printf("%3d. %-20s (%5s) - $%12.2f | Market Cap: $%15.0f | 24h Change: %6.2f%%\n",
			int(crypto["cmc_rank"].(float64)),
			crypto["name"].(string),
			crypto["symbol"].(string),
			usd["price"].(float64),
			usd["market_cap"].(float64),
			usd["percent_change_24h"].(float64),
		)
	}
}
