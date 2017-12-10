package order

// Order new order
/*
 *	https://coincheck.com/documents/exchange/api#order-new
 *	example: rate: 30000, amount: 10, order_type: "buy", pair: "btc_jpy"
 */
type Order struct {
	ID                     int    `json:"id"`
	Rate                   string `json:"rate"`
	Amount                 string `json:"amount"`
	OrderType              string `json:"order_type"`
	Pair                   string `json:"pair"`
	PendingAmount          string `json:"pending_amount"`
	PendingMarketBuyAmount string `json:"pending_marker_buy_amount"`
	StopLossRate           string `json:"stop_loss_rate"`
	CreatedAt              string `json:"created_at"`
	Success                bool   `json:"success"`
	Error                  string `json:"error"`
}
