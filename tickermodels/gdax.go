package tickermodels

import (
	"time"
)

type GDAXTicker struct {
	LastTradeId    int       `json:"trade_id"`
	LastTradePrice string    `json:"price"`
	LastTradeSize  string    `json:"size"`
	Bid            string    `json:"bid"`
	Ask            string    `json:"ask"`
	Volume         string    `json:"volume"`
	Time           time.Time `json:"time"`
}

type GDAXBestBidAsk struct {
	Sequence int64           `json:"sequence"`
	Bids     [][]interface{} `json:"bids"`
	Asks     [][]interface{} `json:"asks"`
}

func (response GDAXTicker) GetExchangeData() Ticker {
	return Ticker{
		Exchange: "GDAX",
		Bid:      response.Bid,
		Ask:      response.Ask,
	}
}

//{
//    "sequence": "3",
//    "bids": [
//        [ price, size, num-orders ],
//        [ "295.96", "4.39088265", 2 ],
//    ],
//    "asks": [
//        [ price, size, num-orders ],
//        [ "295.97", "25.23542881", 12 ],
//    ]
//}

func (response GDAXBestBidAsk) GetExchangeData() Ticker {
	return Ticker{
		Exchange:         "GDAX",
		Bid:              response.Bids[0][0].(string),
		BidSize:          response.Bids[0][1].(string),
		Ask:              response.Asks[0][0].(string),
		AskSize:          response.Asks[0][1].(string),
		DepositFee:       "0.0",
		MakerFee:         "0.0",
		TakerFee:         "0.25",
		BTCWithdrawalFee: "0.0",
		HasWithdrawalFee: false,
	}
}
