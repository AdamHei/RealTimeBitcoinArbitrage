package tickermodels

type ExmoTicker struct {
	BTC_USD exmoInnerTicker `json:"BTC_USD"`
}

type exmoInnerTicker struct {
	AskQuantity string     `json:"ask_quantity"`
	AskAmount   string     `json:"ask_amount"`
	AskTop      string     `json:"ask_top"`
	BidQuantity string     `json:"bid_quantity"`
	BidAmount   string     `json:"bid_amount"`
	BidTop      string     `json:"bid_top"`
	Ask         [][]string `json:"ask"`
	Bid         [][]string `json:"bid"`
}

func (response ExmoTicker) GetExchangeData() Ticker {
	return Ticker{
		Exchange:         "ExMo",
		Bid:              response.BTC_USD.Bid[0][0],
		BidSize:          response.BTC_USD.Bid[0][1],
		Ask:              response.BTC_USD.Ask[0][0],
		AskSize:          response.BTC_USD.Ask[0][1],
		DepositFee:       "0.0",
		MakerFee:         "0.2",
		TakerFee:         "0.2",
		BTCWithdrawalFee: "0.0005",
		HasWithdrawalFee: true,
	}
}
