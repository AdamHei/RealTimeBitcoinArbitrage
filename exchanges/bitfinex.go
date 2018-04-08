package exchanges

import (
	"encoding/json"
	"github.com/adamhei/honorsproject/tickermodels"
	"net/http"
)

const bitfinexEndpoint = "https://api.bitfinex.com/v2/ticker/tBTCUSD"

func fetchBidAskBitfinex(ch chan<- tickermodels.Ticker) {
	resp, err := http.Get(bitfinexEndpoint)
	if err != nil {
		errorHandler("Could not fetch data from Bitfinex:"+err.Error(), ch)
		return
	}

	bitTickerArr := make([]float64, 0)
	err = json.NewDecoder(resp.Body).Decode(&bitTickerArr)
	resp.Body.Close()
	if err != nil {
		errorHandler("Could not parse Bitfinex response"+err.Error(), ch)
		return
	}

	bitResponse := tickermodels.BitfinexV2Ticker{
		Bid:     bitTickerArr[0],
		BidSize: bitTickerArr[1],
		Ask:     bitTickerArr[2],
		AskSize: bitTickerArr[3],
	}

	ch <- bitResponse.GetExchangeData()
}
