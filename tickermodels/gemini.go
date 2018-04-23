package tickermodels

type GeminiBestBidAsk struct {
	Bids []GeminiOpenOrder `json:"bids"`
	Asks []GeminiOpenOrder `json:"asks"`
}

// Either a Bid or an Ask
type GeminiOpenOrder struct {
	Price     string `json:"price"`
	Amount    string `json:"amount"`
	Timestamp string `json:"timestamp"`
}

func (response GeminiBestBidAsk) GetExchangeData() Ticker {
	return Ticker{
		Exchange:         "Gemini",
		Bid:              response.Bids[0].Price,
		BidSize:          response.Bids[0].Amount,
		Ask:              response.Asks[0].Price,
		AskSize:          response.Asks[0].Amount,
		DepositFee:       "0.0",
		MakerFee:         "1.0",
		TakerFee:         "1.0",
		BTCWithdrawalFee: "0.0",
		HasWithdrawalFee: false,
	}
}
