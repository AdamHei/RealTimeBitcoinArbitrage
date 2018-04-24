package exchanges

import (
	"encoding/json"
	"github.com/adamhei/honorsproject/tickermodels"
	"net/http"
)

const gdaxEndpoint = "https://api.gdax.com/products/BTC-USD/book"

func fetchBidAskGDAX(ch chan<- tickermodels.Ticker) {
	resp, err := http.Get(gdaxEndpoint)
	if err != nil {
		errorHandler("Could not fetch data from GDAX:", err, ch)
		return
	}

	gdaxResponse := new(tickermodels.GDAXBestBidAsk)
	err = json.NewDecoder(resp.Body).Decode(gdaxResponse)
	resp.Body.Close()

	if err != nil {
		errorHandler("Could not parse GDAX json:", err, ch)
		return
	}

	ch <- gdaxResponse.GetExchangeData()
}
