package tickermodels

type BinanceTicker struct {
	Symbol      string `json:"symbol"`
	BidPrice    string `json:"bidPrice"`
	BidQuantity string `json:"bidQty"`
	AskPrice    string `json:"askPrice"`
	AskQuantity string `json:"askQty"`
}

func (response BinanceTicker) GetExchangeData() LimitedJson {
	return LimitedJson{
		"Binance": {
			BID: response.BidPrice,
			ASK: response.AskPrice,
		},
	}
}
