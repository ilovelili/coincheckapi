package ticker

// Ticker represents Coincheck ticker
type Ticker struct {
	Ask       float64 `json:"ask"`
	Bid       float64 `json:"bid"`
	High      float64 `json:"high"`
	Last      float64 `json:"last"`
	Low       float64 `json:"low"`
	Timestamp float64 `json:"timestamp"`
	Volume    string  `json:"volume"`
}
