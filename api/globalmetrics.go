package api

import "github.com/tigusigalpa/coinmarketcap-go/http"

type GlobalMetrics struct {
	httpClient *http.Client
}

func NewGlobalMetrics(httpClient *http.Client) *GlobalMetrics {
	return &GlobalMetrics{
		httpClient: httpClient,
	}
}

func (g *GlobalMetrics) QuotesLatest(params map[string]string) (map[string]interface{}, error) {
	return g.httpClient.Get("/v1/global-metrics/quotes/latest", params)
}

func (g *GlobalMetrics) QuotesHistorical(params map[string]string) (map[string]interface{}, error) {
	return g.httpClient.Get("/v1/global-metrics/quotes/historical", params)
}
