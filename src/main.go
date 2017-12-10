package main

import (
	"coincheck"
	"coincheck/balance"
	"coincheck/ticker"
	"config"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// init client
	configure := config.GetConfig()
	client := coincheck.New(configure.Key, configure.SecretKey)

	// step 1. Check balance
	balance, err := balance.GetBalance(client)
	if err != nil {
		panic(err)
	}

	balanceinjpy, err := strconv.ParseFloat(balance.JPY, 32)
	if err != nil {
		panic(err)
	}

	if balanceinjpy < 1000 {
		fmt.Printf("Insufficient balance. Your current balance is %g yen, please depoist\n", balanceinjpy)
		os.Exit(1)
	}

	// Step 2. Check exchange rate
	ticker, err := ticker.GetTicker(client)
	if err != nil {
		panic(err)
	}

	fmt.Println(ticker)

	// Step 3. Apply rules

}
