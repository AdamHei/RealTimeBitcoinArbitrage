package exchanges

import (
	"encoding/json"
	"github.com/adamhei/honorsproject/exchanges/tickermodels"
	"net/http"
)

const geminiEndpoint = "https://api.gemini.com/v1/pubticker/btcusd"

func fetchBidAskData(ch chan<- tickermodels.LimitedJson) {
	resp, err := http.Get(geminiEndpoint)
	if err != nil {
		errorHandler("Could not fetch Gemini data:"+err.Error(), ch)
	}

	geminiResponse := new(tickermodels.GeminiTicker)
	err = json.NewDecoder(resp.Body).Decode(geminiResponse)

	resp.Body.Close()
	if err != nil {
		errorHandler("Could not parse Gemini json:"+err.Error(), ch)
		return
	}

	ch <- geminiResponse.GetExchangeData()
}
