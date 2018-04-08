package exchanges

import "github.com/adamhei/honorsproject/tickermodels"

const NumExchanges = 6

func FetchAllExchanges(ch chan<- tickermodels.Ticker) {
	go fetchBidAskPoloniex(ch)
	go fetchBidAskData(ch)
	go fetchBidAskKraken(ch)
	go fetchBidAskGDAX(ch)
	go fetchBidAskBitfinex(ch)
	go fetchBidAskBinance(ch)
}
