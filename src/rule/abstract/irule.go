package rule

import "coincheck/ticker"

// IRule rule interface
type IRule interface {
	Apply(*ticker.ExchangeRate)
}
