package balance

import (
	"coincheck"
	"errors"
)

// GetBalance returns account balance
func GetBalance(api *coincheck.APIClient) (balance Balance, err error) {
	err = api.DoGetRequest(BalanceEndpointURL, []byte(""), &balance)
	if err != nil {
		return balance, err
	}
	if !balance.Success {
		return balance, errors.New(balance.Error)
	}
	return balance, nil
}
