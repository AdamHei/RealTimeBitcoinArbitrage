package exchanges

import (
	"encoding/json"
	"github.com/adamhei/honorsproject/tickermodels"
	"net/http"
	"time"
)

const poloniexEndpoint = "http://poloniex.com/public?command=returnTicker"
const USDT_BTC = "USDT_BTC"

func fetchBidAskPoloniex(ch chan<- tickermodels.Ticker) {
	// Poloniex can be slow, so we nix the request if it takes too long
	timeout := time.Duration(500 * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(poloniexEndpoint)
	if err != nil {
		errorHandler("Could not fetch Poloniex data: "+err.Error(), ch)
		return
	}

	fullResponse := new(map[string]tickermodels.PoloniexTicker)
	err = json.NewDecoder(resp.Body).Decode(fullResponse)
	resp.Body.Close()

	if err != nil {
		errorHandler("Could not parse Poloniex json: "+err.Error(), ch)
		return
	}

	btcTicker := (*fullResponse)[USDT_BTC]
	ch <- btcTicker.GetExchangeData()
}
