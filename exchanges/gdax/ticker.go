package gdax

import (
	"encoding/json"
	"github.com/adamhei/honorsproject/exchanges/errorhandling"
	"github.com/adamhei/honorsproject/exchanges/tickermodels"
	"net/http"
)

const tickerUrl = "https://api.gdax.com/products/BTC-USD/ticker"

func FetchBidAskGDAX(ch chan<- tickermodels.LimitedJson) {
	resp, err := http.Get(tickerUrl)
	if err != nil {
		errorhandling.ErrorHandler("Could not fetch data from GDAX:"+err.Error(), ch)
		return
	}

	gdaxResponse := new(tickermodels.GDAXTicker)
	err = json.NewDecoder(resp.Body).Decode(gdaxResponse)
	resp.Body.Close()

	if err != nil {
		errorhandling.ErrorHandler("Could not parse GDAX json:"+err.Error(), ch)
		return
	}

	ch <- gdaxResponse.GetExchangeData()
}
