package exchanges

import (
	"encoding/json"
	"github.com/adamhei/honorsproject/tickermodels"
	"net/http"
)

//const geminiEndpoint = "https://api.gemini.com/v1/pubticker/btcusd"
const geminiEndpoint = "https://api.gemini.com/v1/book/btcusd?limit_bids=1&limit_asks=1"

func fetchBidAskGemini(ch chan<- tickermodels.Ticker) {
	resp, err := http.Get(geminiEndpoint)
	if err != nil {
		errorHandler("Could not fetch Gemini data:", err, ch)
	}

	geminiResponse := new(tickermodels.GeminiBestBidAsk)
	err = json.NewDecoder(resp.Body).Decode(geminiResponse)

	resp.Body.Close()
	if err != nil {
		errorHandler("Could not parse Gemini json:", err, ch)
		return
	}

	ch <- geminiResponse.GetExchangeData()
}
