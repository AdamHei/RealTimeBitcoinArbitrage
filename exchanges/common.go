package exchanges

import "github.com/adamhei/honorsproject/exchanges/tickermodels"

const NUMEXCHANGES = 6

func FetchAllExchanges(ch chan<- tickermodels.LimitedJson) {
	go fetchBidAskPoloniex(ch)
	go fetchBidAskData(ch)
	go fetchBidAskKraken(ch)
	go fetchBidAskGDAX(ch)
	go fetchBidAskBitfinex(ch)
	go fetchBidAskBinance(ch)
}
