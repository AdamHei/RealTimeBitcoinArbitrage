package exchanges

import (
	"github.com/adamhei/honorsproject/exchanges/bitfinex"
	"github.com/adamhei/honorsproject/exchanges/gdax"
	"github.com/adamhei/honorsproject/exchanges/gemini"
	"github.com/adamhei/honorsproject/exchanges/kraken"
	"github.com/adamhei/honorsproject/exchanges/models"
	"github.com/adamhei/honorsproject/exchanges/poloniex"
)

const NUMEXCHANGES = 5

func FetchAllExchanges(ch chan<- models.LimitedJson) {
	go poloniex.FetchBidAskPoloniex(ch)
	go gemini.FetchBidAskData(ch)
	go kraken.FetchBidAskKraken(ch)
	go gdax.FetchBidAskGDAX(ch)
	go bitfinex.FetchBidAskBitfinex(ch)
}
