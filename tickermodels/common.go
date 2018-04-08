package tickermodels

import (
	"fmt"
	"strconv"
)

type ITicker interface {
	GetExchangeData() Ticker
}

type Ticker struct {
	Exchange string `json:"exchange"`
	Bid      string `json:"bid"`
	BidSize  string `json:"bid_size"`
	Ask      string `json:"ask"`
	AskSize  string `json:"ask_size"`
	// Percentages
	DepositFee string `json:"deposit_fee"`
	MakerFee   string `json:"maker_fee"`
	TakerFee   string `json:"taker_fee"`
	// In BTC
	BTCWithdrawalFee string `json:"btc_withdrawal_fee"`
}

func (ticker Ticker) String() string {
	return fmt.Sprintf("%v\nBid: %v\nBidSize: %v\nAsk: %v\nAskSize: %v\n", ticker.Exchange, ticker.Bid, ticker.BidSize, ticker.Ask, ticker.AskSize)
}

//buy exchange
//ask
//ask size
//taker fee
//btc withdrawal fee
//
//sell exchange
//bid
//bid size
//taker fee
//
//// TODO BTC network fee and est transfer time

func basicFormatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
