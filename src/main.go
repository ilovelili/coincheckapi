package main

import (
	"coincheck/account"
	"config"

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
	message := util.Nonce
	signature := resolveSignature(secretKey, message)

	err = account.GetAccountBalance(publicKey, util.Nonce, signature)

	if err != nil {
		panic(err)
	}
}

// resolveSignature resolve signature
// https://coincheck.com/documents/exchange/api
func resolveSignature(secret, message string) string {
	secretbytes := []byte(secret)
	messagebytes := []byte(message)

	return util.HmacSHA256(secretbytes, messagebytes)
}
