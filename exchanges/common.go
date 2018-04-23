// Package exchanges interfaces directly with exchange APIs to pull real-time bid/ask spreads
//
// Each file corresponds to an exchange, which must implement a function of the form
// `fetchBidAskX(ch chan<-Ticker)` and pass its exchange's Ticker
package exchanges

import "github.com/adamhei/honorsproject/tickermodels"

// The number of supported exchanges
const NumExchanges = 7

// FetchAllExchanges calls every exchange's Ticker-fetching method
func FetchAllExchanges(ch chan<- tickermodels.Ticker) {
	go fetchBidAskPoloniex(ch)
	go fetchBidAskGemini(ch)
	go fetchBidAskKraken(ch)
	go fetchBidAskGDAX(ch)
	go fetchBidAskBitfinex(ch)
	go fetchBidAskBinance(ch)
	go fetchBidAskBitstamp(ch)
}
