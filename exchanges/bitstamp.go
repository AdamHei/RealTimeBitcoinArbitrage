package exchanges

import (
	"encoding/json"
	"github.com/adamhei/honorsproject/tickermodels"
	"net/http"
)

const bitstampUrl = "https://www.bitstamp.net/api/order_book/"

func fetchBidAskBitstamp(ch chan<- tickermodels.Ticker) {
	resp, err := http.Get(bitstampUrl)
	if err != nil {
		errorHandler("Could not fetch order book from Bitstamp"+err.Error(), ch)
		return
	}

	bitstampTicker := new(tickermodels.BitstampTicker)
	err = json.NewDecoder(resp.Body).Decode(bitstampTicker)
	resp.Body.Close()
	if err != nil {
		errorHandler("Could not parse Bitstamp response"+err.Error(), ch)
		return
	}

	ch <- bitstampTicker.GetExchangeData()
}
