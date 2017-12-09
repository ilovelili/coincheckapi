package util

import (
	"strconv"
	"time"
)

var (
	// Nonce unix time nonce in string
	Nonce = strconv.FormatInt(time.Now().Unix(), 10)
)
