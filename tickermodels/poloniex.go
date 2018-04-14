package tickermodels

type PoloniexTicker struct {
	Id            int    `json:"id"`
	LastPrice     string `json:"last"`
	Ask           string `json:"lowestAsk"`
	Bid           string `json:"highestBid"`
	PercentChange string `json:"percentChange"`
	BaseVolume    string `json:"baseVolume"`
	QuoteVolume   string `json:"quoteVolume"`
	IsFrozen      string `json:"isFrozen"`
	High24Hr      string `json:"high24hr"`
	Low24Hr       string `json:"low24hr"`
}

type PoloniexBestBidAsk struct {
	Sequence int64           `json:"seq"`
	IsFrozen string          `json:"isFrozen"`
	Asks     [][]interface{} `json:"asks"`
	Bids     [][]interface{} `json:"bids"`
}

func (response PoloniexTicker) GetExchangeData() Ticker {
	return Ticker{
		Exchange: "Poloniex",
		Bid:      response.Bid,
		Ask:      response.Ask,
	}
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
