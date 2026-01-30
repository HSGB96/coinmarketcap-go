package api

import "github.com/tigusigalpa/coinmarketcap-go/http"

type Cryptocurrency struct {
	httpClient *http.Client
}

func NewCryptocurrency(httpClient *http.Client) *Cryptocurrency {
	return &Cryptocurrency{
		httpClient: httpClient,
	}
}

func (c *Cryptocurrency) ListingsLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/listings/latest", params)
}

func (c *Cryptocurrency) ListingsHistorical(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/listings/historical", params)
}

func (c *Cryptocurrency) QuotesLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/quotes/latest", params)
}

func (c *Cryptocurrency) QuotesHistorical(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/quotes/historical", params)
}

func (c *Cryptocurrency) Info(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/info", params)
}

func (c *Cryptocurrency) Map(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/map", params)
}

func (c *Cryptocurrency) OHLCVLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/ohlcv/latest", params)
}

func (c *Cryptocurrency) OHLCVHistorical(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v2/cryptocurrency/ohlcv/historical", params)
}

func (c *Cryptocurrency) Categories(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/categories", params)
}

func (c *Cryptocurrency) TrendingLatest(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/trending/latest", params)
}

func (c *Cryptocurrency) TrendingGainersLosers(params map[string]string) (map[string]interface{}, error) {
	return c.httpClient.Get("/v1/cryptocurrency/trending/gainers-losers", params)
}
