package rule

// AskRule ask rule
type AskRule struct {
	Price float32
	Threshold float32
	DailyBudget float32
}

// Apply apply the rule and sell bitcoin
func (rule *AskRule) Apply(exchangeRate *ticker.ExchangeRate) {
	
}