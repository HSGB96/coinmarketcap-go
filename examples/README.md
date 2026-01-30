# CoinMarketCap Go SDK Examples

This directory contains various examples demonstrating how to use the CoinMarketCap Go SDK.

## Available Examples

### 1. Basic Example (`basic/`)
Simple example showing how to get Bitcoin and Ethereum prices.

```bash
cd basic
go run main.go
```

### 2. Top 100 Cryptocurrencies (`top100/`)
Fetches and displays the top 100 cryptocurrencies by market cap.

```bash
cd top100
go run main.go
```

### 3. Error Handling (`error_handling/`)
Demonstrates comprehensive error handling for all API error types.

```bash
cd error_handling
go run main.go
```

### 4. Global Metrics (`global_metrics/`)
Shows how to fetch global cryptocurrency market metrics.

```bash
cd global_metrics
go run main.go
```

### 5. With Environment Variables (`with_env/`)
Example using environment variables from a `.env` file.

```bash
cd with_env
# Create .env file with your API key
echo "COINMARKETCAP_API_KEY=your-api-key-here" > .env
go run main.go
```

## Before Running Examples

Make sure to replace `your-api-key-here` with your actual CoinMarketCap API key in each example, or use environment variables.

### Get Your API Key

1. Visit [CoinMarketCap API Portal](https://coinmarketcap.com/api/)
2. Sign up for a free account
3. Get your API key from the dashboard

## Running All Examples

You can use the Makefile in the root directory:

```bash
# From the root directory
make run-basic
make run-top100
make run-error-handling
make run-global-metrics
make run-with-env
```

## Additional Resources

- [Main Documentation](../README.md)
- [CoinMarketCap API Documentation](https://coinmarketcap.com/api/documentation/v1/)
