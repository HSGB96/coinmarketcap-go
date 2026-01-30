# Release Notes

## Version 1.0.0 - Initial Release (2024-01-30)

### 🎉 First Stable Release

We're excited to announce the first stable release of the CoinMarketCap Go SDK! This professional-grade client library provides complete integration with the CoinMarketCap API v1 for Golang applications.

### ✨ Features

#### Complete API Coverage
- **Cryptocurrency API** - 11 endpoints covering listings, quotes, metadata, OHLCV data, categories, and trending cryptocurrencies
- **Exchange API** - 6 endpoints for exchange listings, quotes, metadata, and market pairs
- **Global Metrics API** - 2 endpoints for current and historical global market metrics
- **Tools API** - Price conversion endpoint for multi-currency conversions

#### Developer Experience
- **Fluent Builder Pattern** - Intuitive, chainable API for easy client configuration
- **Type-Safe Error Handling** - Comprehensive error types with proper Go error wrapping
  - `AuthenticationError` (401) - Invalid or missing API key
  - `RateLimitError` (429) - Rate limit exceeded with retry-after support
  - `InvalidRequestError` (400) - Invalid request parameters
  - `NotFoundError` (404) - Resource not found
  - `APIError` - Generic API errors with full context
- **Sandbox Mode Support** - Test your integration without consuming production API credits
- **Configurable Timeouts** - Customize HTTP request timeouts to match your needs
- **Environment Variable Support** - Easy configuration via `.env` files using godotenv

#### Documentation & Examples
- **Comprehensive README** - Available in English and Russian
- **5 Working Examples**:
  - Basic usage - Get Bitcoin and Ethereum prices
  - Top 100 cryptocurrencies - Market cap rankings
  - Error handling - Comprehensive error handling patterns
  - Global metrics - Market-wide statistics
  - Environment variables - Configuration via .env files
- **API Reference Tables** - Quick lookup for all available methods
- **Contributing Guidelines** - Clear instructions for contributors
- **MIT License** - Free to use in commercial and open-source projects

#### Code Quality
- **Go 1.21+ Support** - Modern Golang features and best practices
- **Clean Architecture** - Well-organized package structure with clear separation of concerns
- **Idiomatic Go Code** - Follows Go conventions and community standards
- **Production Ready** - Battle-tested error handling and edge cases covered

### 📦 Installation

```bash
go get github.com/tigusigalpa/coinmarketcap-go
```

### 🚀 Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    coinmarketcap "github.com/tigusigalpa/coinmarketcap-go"
)

func main() {
    client, err := coinmarketcap.NewClientBuilder().
        SetAPIKey("your-api-key").
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
    
    fmt.Printf("Quotes: %v\n", quotes)
}
```

### 📋 What's Included

**Core Package:**
- `client.go` - Main client and fluent builder
- `errors/errors.go` - Type-safe error definitions
- `http/client.go` - HTTP client implementation
- `api/cryptocurrency.go` - Cryptocurrency endpoints
- `api/exchange.go` - Exchange endpoints
- `api/globalmetrics.go` - Global metrics endpoints
- `api/tools.go` - Tools endpoints

**Documentation:**
- `README.md` - English documentation
- `README-ru.md` - Russian documentation
- `CHANGELOG.md` - Version history
- `CONTRIBUTING.md` - Contribution guidelines
- `LICENSE` - MIT License

**Development Tools:**
- `Makefile` - Build automation and common tasks
- `.editorconfig` - Editor configuration
- `.gitignore` - Git ignore rules
- `go.mod` & `go.sum` - Go module dependencies

**Examples:**
- `examples/basic/` - Basic usage example
- `examples/top100/` - Top 100 cryptocurrencies
- `examples/error_handling/` - Error handling patterns
- `examples/global_metrics/` - Global market metrics
- `examples/with_env/` - Environment variables configuration

### 🔗 Resources

- **GitHub Repository**: https://github.com/tigusigalpa/coinmarketcap-go
- **CoinMarketCap API Documentation**: https://coinmarketcap.com/api/documentation/v1/
- **Get API Key**: https://coinmarketcap.com/api/

### 🙏 Acknowledgments

This project is a Golang port of the CoinMarketCap PHP library, maintaining the same architecture and API design while following Go idioms and best practices.

### 📝 License

MIT License - Free to use in commercial and open-source projects.

---

**Full Changelog**: https://github.com/tigusigalpa/coinmarketcap-go/commits/v1.0.0
