// Package tickermodels contains exchange-specific data models representing response bodies and a unified data model: Ticker
//
// Each file contains the exchange-specific data model and an implementation of GetExchangeData
package tickermodels

import (
	"fmt"
	"strconv"
)

// Every exchange Ticker model must conform to this interface
type ITicker interface {
	// GetExchangeData returns the exchange's information in Ticker format
	GetExchangeData() Ticker
}

// Unified data model containing every property relevant to the spread on an exchange and its fees
type Ticker struct {
	Exchange string `json:"exchange"`
	Bid      string `json:"bid"`
	BidSize  string `json:"bid_size"`
	Ask      string `json:"ask"`
	AskSize  string `json:"ask_size"`
	// Percentages in decimal form, i.e. .25% = 0.25
	DepositFee string `json:"deposit_fee"`
	MakerFee   string `json:"maker_fee"`
	TakerFee   string `json:"taker_fee"`
	// In BTC
	BTCWithdrawalFee string `json:"btc_withdrawal_fee"`
	HasWithdrawalFee bool `json:"has_withdrawal_fee"`
}

func (ticker Ticker) String() string {
	return fmt.Sprintf("%v\nBid: %v\nBidSize: %v\nAsk: %v\nAskSize: %v\n", ticker.Exchange, ticker.Bid, ticker.BidSize, ticker.Ask, ticker.AskSize)
}

func basicFormatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
