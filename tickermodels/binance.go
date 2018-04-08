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
		Exchange:         "Binance",
		Bid:              response.BidPrice,
		BidSize:          response.BidQuantity,
		Ask:              response.AskPrice,
		AskSize:          response.AskQuantity,
		DepositFee:       "0.0",
		MakerFee:         "0.1",
		TakerFee:         "0.1",
		BTCWithdrawalFee: "0.0005",
	}
}
