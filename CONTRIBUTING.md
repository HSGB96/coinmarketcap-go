# Contributing to CoinMarketCap Go SDK

Thank you for your interest in contributing to the CoinMarketCap Go SDK! This document provides guidelines and instructions for contributing.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/your-username/coinmarketcap-go.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Run tests: `go test ./...`
6. Commit your changes: `git commit -am 'Add some feature'`
7. Push to the branch: `git push origin feature/your-feature-name`
8. Submit a pull request

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git
- A CoinMarketCap API key for testing

### Installation

```bash
# Clone the repository
git clone https://github.com/tigusigalpa/coinmarketcap-go.git
cd coinmarketcap-go

# Install dependencies
go mod download

# Run tests
go test ./...
```

## Code Style

- Follow standard Go conventions and idioms
- Use `gofmt` to format your code
- Use `golint` to check for style issues
- Write clear, descriptive commit messages
- Add comments for exported functions and types

## Testing

- Write tests for new features
- Ensure all tests pass before submitting a PR
- Aim for high test coverage
- Use table-driven tests where appropriate

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

## Pull Request Process

1. Update the README.md with details of changes if applicable
2. Update the CHANGELOG.md with your changes
3. Ensure all tests pass
4. Update documentation if you're changing functionality
5. The PR will be merged once you have the sign-off of at least one maintainer

## Code of Conduct

### Our Pledge

We are committed to providing a welcoming and inspiring community for all.

### Our Standards

- Be respectful and inclusive
- Be patient and welcoming
- Be collaborative
- Be mindful of your words

## Questions?

Feel free to open an issue for any questions or concerns.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
