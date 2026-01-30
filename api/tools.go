package api

import "github.com/tigusigalpa/coinmarketcap-go/http"

type Tools struct {
	httpClient *http.Client
}

func NewTools(httpClient *http.Client) *Tools {
	return &Tools{
		httpClient: httpClient,
	}
}

func (t *Tools) PriceConversion(params map[string]string) (map[string]interface{}, error) {
	return t.httpClient.Get("/v2/tools/price-conversion", params)
}
