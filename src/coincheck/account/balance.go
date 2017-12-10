package account

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// GetAccountBalance get account balance
func GetAccountBalance(key, nonce, signature string) (balance *BalanceResponse, err error) {
	req, _ := http.NewRequest("GET", BalanceEndpointUrl, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	req.Header.Set("ACCESS-KEY", key)
	req.Header.Set("ACCESS-NONCE", nonce)
	req.Header.Set("ACCESS-SIGNATURE", signature)

	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		bytes, _ := ioutil.ReadAll(res.Body)
		err = errors.New(string(bytes))
		return
	}

	err = json.NewDecoder(res.Body).Decode(&balance)
	return
}
