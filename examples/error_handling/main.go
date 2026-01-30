package main

import (
	"errors"
	"fmt"
	"log"

	coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
	cmcerrors "github.com/tigusigalpa/coinmarketcap-go/errors"
)

func main() {
	client, err := coinmarketcap.NewClientBuilder().
		SetAPIKey("your-api-key-here").
		Build()

	if err != nil {
		log.Fatal(err)
	}

	listings, err := client.Cryptocurrency().ListingsLatest(map[string]string{
		"limit":   "100",
		"convert": "USD",
	})

	if err != nil {
		var authErr *cmcerrors.AuthenticationError
		var rateLimitErr *cmcerrors.RateLimitError
		var invalidReqErr *cmcerrors.InvalidRequestError
		var notFoundErr *cmcerrors.NotFoundError
		var apiErr *cmcerrors.APIError

		switch {
		case errors.As(err, &authErr):
			fmt.Printf("❌ Authentication failed: %v\n", authErr)
			fmt.Println("Please check your API key in the .env file")
		case errors.As(err, &rateLimitErr):
			fmt.Printf("⏱️  Rate limit exceeded: %v\n", rateLimitErr)
			if rateLimitErr.GetRetryAfter() != nil {
				fmt.Printf("Retry after: %d seconds\n", *rateLimitErr.GetRetryAfter())
			}
		case errors.As(err, &invalidReqErr):
			fmt.Printf("⚠️  Invalid request: %v\n", invalidReqErr)
			fmt.Println("Please check your request parameters")
		case errors.As(err, &notFoundErr):
			fmt.Printf("🔍 Resource not found: %v\n", notFoundErr)
		case errors.As(err, &apiErr):
			fmt.Printf("🚨 API error [%d]: %v\n", apiErr.StatusCode, apiErr)
		default:
			fmt.Printf("❓ Unknown error: %v\n", err)
		}
		return
	}

	fmt.Printf("✅ Successfully retrieved listings: %v\n", listings)
}
