package ticker

import "coincheck"

// GetTicker returns Coincheck ticker
func GetTicker(api *coincheck.APIClient) (ticker Ticker, err error) {
	err = api.DoGetRequest(TickerEndpoint, []byte(""), &ticker)
	if err != nil {
		return ticker, err
	}
	return ticker, nil
}
