package tickermodels

type KrakenTicker struct {
	Error  []string `json:"error"`
	Result result   `json:"result"`
}

type result struct {
	XBTUSD pairBody `json:"XXBTZUSD"`
}

type pairBody struct {
	AskArr               []string `json:"a"`
	BidArr               []string `json:"b"`
	ClosedTradeArr       []string `json:"c"`
	VolumeArr            []string `json:"v"`
	VolumeWeightedAvgArr []string `json:"p"`
	NumTradesArr         []int    `json:"t"`
	LowArr               []string `json:"l"`
	HighArr              []string `json:"h"`
	OpenPrice            string   `json:"o"`
}

type KrakenBestBidAsk struct {
	Error  []string     `json:"error"`
	Result krakenResult `json:"result"`
}

type krakenResult struct {
	TopOrders krakenTopOrders `json:"XXBTZUSD"`
}

type krakenTopOrders struct {
	Asks [][]interface{} `json:"asks"`
	Bids [][]interface{} `json:"bids"`
}

func (response KrakenTicker) GetExchangeData() Ticker {
	return Ticker{
		Exchange: "Kraken",
		Bid:      response.Result.XBTUSD.BidArr[0],
		Ask:      response.Result.XBTUSD.AskArr[0],
	}
}

func (response KrakenBestBidAsk) GetExchangeData() Ticker {
	return Ticker{
		Exchange: "Kraken",
		Bid:      response.Result.TopOrders.Bids[0][0].(string),
		BidSize:  response.Result.TopOrders.Bids[0][1].(string),
		Ask:      response.Result.TopOrders.Asks[0][0].(string),
		AskSize:  response.Result.TopOrders.Asks[0][1].(string),
		// Site is down; assumed to be 0
		DepositFee:       "0.0",
		MakerFee:         "0.16",
		TakerFee:         "0.26",
		BTCWithdrawalFee: "0.0005",
	}
}
