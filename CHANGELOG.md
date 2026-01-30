# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2024-01-30

### Added
- Initial release of CoinMarketCap Go SDK
- Full support for CoinMarketCap API v1
- Cryptocurrency API endpoints
  - Listings (latest and historical)
  - Quotes (latest and historical)
  - Info and metadata
  - Map for ID mapping
  - OHLCV data (latest and historical)
  - Categories
  - Trending cryptocurrencies
  - Gainers and losers
- Exchange API endpoints
  - Listings
  - Quotes (latest and historical)
  - Info and metadata
  - Map for ID mapping
  - Market pairs
- Global Metrics API endpoints
  - Latest quotes
  - Historical quotes
- Tools API endpoints
  - Price conversion
- Comprehensive error handling
  - AuthenticationError
  - RateLimitError
  - InvalidRequestError
  - NotFoundError
  - APIError
- ClientBuilder with fluent interface
- Sandbox mode support
- Configurable timeout
- Environment variable support
- Complete documentation
- Usage examples
- MIT License

### Features
- Type-safe error handling
- Clean and idiomatic Go code
- Easy to use fluent API
- Production-ready
- Well-documented with examples
