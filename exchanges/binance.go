package exchanges

import (
	"encoding/json"
	"fmt"
	"github.com/adamhei/honorsproject/tickermodels"
	"net/http"
	"strconv"
)

const BTCUSDT = "BTCUSDT"
const binanceBaseUrl = "https://api.binance.com"
const binanceApiVersion = "v3"
var binanceEndpoint = fmt.Sprintf("%s/api/%s/ticker/bookTicker", binanceBaseUrl, binanceApiVersion)

func fetchBidAskBinance(ch chan<- tickermodels.Ticker) {
	binanceRequest := buildBinanceRequest()
	resp, err := http.Get(binanceRequest)
	if err != nil {
		errorHandler("Could not fetch data from Bitfinex", err, ch)
		return
	}

	binanceTicker := new(tickermodels.BinanceTicker)
	err = json.NewDecoder(resp.Body).Decode(binanceTicker)
	resp.Body.Close()
	if err != nil {
		errorHandler("Could not parse Binance response", err, ch)
		return
	}

	// Binance gives way too many zeros after the second decimal digit
	// This removes them
	askPrice, _ := strconv.ParseFloat(binanceTicker.AskPrice, 64)
	binanceTicker.AskPrice = strconv.FormatFloat(askPrice, 'f', 2, 64)
	bidPrice, _ := strconv.ParseFloat(binanceTicker.BidPrice, 64)
	binanceTicker.BidPrice = strconv.FormatFloat(bidPrice, 'f', 2, 64)

	ch <- binanceTicker.GetExchangeData()
}

func buildBinanceRequest() string {
	request, _ := http.NewRequest(http.MethodGet, binanceEndpoint, nil)
	query := request.URL.Query()
	query.Add("symbol", BTCUSDT)
	request.URL.RawQuery = query.Encode()
	return request.URL.String()
}
