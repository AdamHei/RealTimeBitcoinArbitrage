package tickermodels

type BitstampTicker struct {
	Timestamp string     `json:"timestamp"`
	Bids      [][]string `json:"bids"`
	Asks      [][]string `json:"asks"`
}

func (response BitstampTicker) GetExchangeData() Ticker {
	return Ticker{
		Exchange:         "Bitstamp",
		Bid:              response.Bids[0][0],
		BidSize:          response.Bids[0][1],
		Ask:              response.Asks[0][0],
		AskSize:          response.Asks[0][1],
		DepositFee:       "0.05",
		MakerFee:         "0.25",
		TakerFee:         "0.25",
		BTCWithdrawalFee: "0.0",
		HasWithdrawalFee: false,
	}
}
