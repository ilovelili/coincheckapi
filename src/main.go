package main

import (
	"coincheck/account"
	"config"
	"fmt"

	"util"
)

var (
	publicKey string
	secretKey string
)

func main() {
	configure, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	publicKey = configure.Key
	secretKey = configure.SecretKey
	message := util.Nonce + account.BalanceEndpointUrl
	signature := resolveSignature(secretKey, message)

	if balance, err := account.GetAccountBalance(publicKey, util.Nonce, signature); err != nil {
		panic(err)
	} else {
		fmt.Println(balance)
	}
}

// resolveSignature resolve signature
// https://coincheck.com/documents/exchange/api
func resolveSignature(secret, message string) string {
	secretbytes := []byte(secret)
	messagebytes := []byte(message)

	return util.HmacSHA256(secretbytes, messagebytes)
}
