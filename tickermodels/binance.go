package tickermodels

type BinanceTicker struct {
	Symbol      string `json:"symbol"`
	BidPrice    string `json:"bidPrice"`
	BidQuantity string `json:"bidQty"`
	AskPrice    string `json:"askPrice"`
	AskQuantity string `json:"askQty"`
}

func (response BinanceTicker) GetExchangeData() Ticker {
	return Ticker{
		Bid:      response.BidPrice,
		Ask:      response.AskPrice,
		Exchange: "Binance",
	}
}
