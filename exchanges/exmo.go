package exchanges

import (
	"encoding/json"
	"github.com/adamhei/honorsproject/tickermodels"
	"net/http"
)

const exmoEndpoint = "https://api.exmo.com/v1/order_book/?pair=BTC_USD"

func fetchBidAskExMo(ch chan<- tickermodels.Ticker) {
	resp, err := http.Get(exmoEndpoint)
	if err != nil {
		errorHandler("Could not reach ExMo", err, ch)
		return
	}

	defer resp.Body.Close()

	exmoTicker := new(tickermodels.ExmoTicker)
	err = json.NewDecoder(resp.Body).Decode(exmoTicker)
	if err != nil {
		errorHandler("Could not parse ExMo response", err, ch)
		return
	}

	ch <- exmoTicker.GetExchangeData()
}
