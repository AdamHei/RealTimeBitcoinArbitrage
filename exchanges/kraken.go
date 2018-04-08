package exchanges

import (
	"encoding/json"
	"github.com/adamhei/honorsproject/tickermodels"
	"net/http"
)

const krakenEndpoint = "https://api.kraken.com/0/public/Depth?pair=XBTUSD&count=1"

func fetchBidAskKraken(ch chan<- tickermodels.Ticker) {
	resp, err := http.Get(krakenEndpoint)
	if err != nil {
		errorHandler("Could not fetch Kraken data:"+err.Error(), ch)
	}

	krakenResponse := new(tickermodels.KrakenBestBidAsk)
	err = json.NewDecoder(resp.Body).Decode(krakenResponse)

	resp.Body.Close()
	if err != nil {
		errorHandler("Could not parse Kraken json:"+err.Error(), ch)
		return
	}

	ch <- krakenResponse.GetExchangeData()
}
