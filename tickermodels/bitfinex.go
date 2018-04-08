package tickermodels

type BitfinexTicker struct {
	Mid       string `json:"mid"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	LastPrice string `json:"last_price"`
	Low       string `json:"low"`
	High      string `json:"high"`
	Volume    string `json:"volume"`
	Timestamp string `json:"timestamp"`
}

type BitfinexV2Ticker struct {
	Bid                   float64
	BidSize               float64
	Ask                   float64
	AskSize               float64
	DailyChange           float64
	DailyChangePercentage float64
	LastPrice             float64
	Volume                float64
	High                  float64
	Low                   float64
}

func (ticker BitfinexV2Ticker) GetExchangeData() Ticker {
	return Ticker{
		Exchange:         "Bitfinex",
		Bid:              basicFormatFloat(ticker.Bid),
		BidSize:          basicFormatFloat(ticker.BidSize),
		Ask:              basicFormatFloat(ticker.Ask),
		AskSize:          basicFormatFloat(ticker.AskSize),
		DepositFee:       "0.0",
		MakerFee:         "0.1",
		TakerFee:         "0.2",
		BTCWithdrawalFee: "0.0004",
	}
}

func (response BitfinexTicker) GetExchangeData() Ticker {
	return Ticker{
		Exchange: "Bitfinex",
		Bid:      response.Bid,
		Ask:      response.Ask,
	}
}
