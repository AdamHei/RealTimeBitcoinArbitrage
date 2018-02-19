package exchanges

import (
	"github.com/adamhei/honorsproject/tickermodels"
	"net/http"
	"encoding/json"
)

const krakenEndpoint = "https://api.kraken.com/0/public/Ticker?pair=XBTUSD"

func fetchBidAskKraken(ch chan<- tickermodels.Ticker) {
	resp, err := http.Get(krakenEndpoint)
	if err != nil {
		errorHandler("Could not fetch Kraken data:"+err.Error(), ch)
	}

	krakenResponse := new(tickermodels.KrakenTicker)
	err = json.NewDecoder(resp.Body).Decode(krakenResponse)

	resp.Body.Close()
	if err != nil {
		errorHandler("Could not parse Kraken json:"+err.Error(), ch)
		return
	}

	ch <- krakenResponse.GetExchangeData()
}
