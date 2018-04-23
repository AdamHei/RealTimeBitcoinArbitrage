package tickermodels

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
		HasWithdrawalFee: true,
	}
}
