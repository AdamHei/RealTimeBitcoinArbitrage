package tickermodels

type PoloniexBestBidAsk struct {
	Sequence int64           `json:"seq"`
	IsFrozen string          `json:"isFrozen"`
	Asks     [][]interface{} `json:"asks"`
	Bids     [][]interface{} `json:"bids"`
}

func (response PoloniexBestBidAsk) GetExchangeData() Ticker {
	return Ticker{
		Exchange:         "Poloniex",
		Bid:              response.Bids[0][0].(string),
		BidSize:          basicFormatFloat(response.Bids[0][1].(float64)),
		Ask:              response.Asks[0][0].(string),
		AskSize:          basicFormatFloat(response.Asks[0][1].(float64)),
		DepositFee:       "0.0",
		MakerFee:         "0.15",
		TakerFee:         "0.25",
		BTCWithdrawalFee: "0.0",
		HasWithdrawalFee: false,
	}
}
