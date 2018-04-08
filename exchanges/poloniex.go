package exchanges

import (
	"encoding/json"
	"github.com/adamhei/honorsproject/tickermodels"
	"net/http"
	"time"
)

const poloniexEndpoint = "https://poloniex.com/public?command=returnOrderBook&currencyPair=USDT_BTC&depth=1"

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

	poloniexTicker := new(tickermodels.PoloniexBestBidAsk)
	err = json.NewDecoder(resp.Body).Decode(poloniexTicker)
	resp.Body.Close()

	if err != nil {
		errorHandler("Could not parse Poloniex json: "+err.Error(), ch)
		return
	}

	ch <- poloniexTicker.GetExchangeData()
}
