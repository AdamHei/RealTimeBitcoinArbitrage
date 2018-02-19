package exchanges

import (
	"encoding/json"
	"github.com/adamhei/honorsproject/tickermodels"
	"net/http"
)

const bitfinexEndpoint = "https://api.bitfinex.com/v1/pubticker/btcusd"

func fetchBidAskBitfinex(ch chan<- tickermodels.Ticker) {
	resp, err := http.Get(bitfinexEndpoint)
	if err != nil {
		errorHandler("Could not fetch data from Bitfinex:"+err.Error(), ch)
		return
	}

	bitResponse := new(tickermodels.BitfinexTicker)
	err = json.NewDecoder(resp.Body).Decode(bitResponse)

	resp.Body.Close()
	if err != nil {
		errorHandler("Could not parse Bitfinex response"+err.Error(), ch)
		return
	}

	ch <- bitResponse.GetExchangeData()
}
