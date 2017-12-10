package balance

// Balance balance response
type Balance struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	JPY     string `json:"jpy"`
	BTC     string `json:"btc"`
	ETH     string `json:"eth"`
	ETC     string `json:"etc"`
	DAO     string `json:"dao"`
	LSK     string `json:"lsk"`
	FCT     string `json:"fct"`
	XMR     string `json:"xmr"`
	REP     string `json:"rep"`
	XRP     string `json:"xrp"`
	ZEC     string `json:"zec"`
	XEM     string `json:"xem"`
	LTC     string `json:"ltc"`
	DASH    string `json:"dash"`
	BCH     string `json:"bch"`
}
