package handlers

import (
	"github.com/adamhei/honorsproject/exchanges"
	"github.com/adamhei/honorsproject/exchanges/tickermodels"
	"log"
	"math"
	"net/http"
	"strconv"
)

// Return the biggest spread between the highest bid and the lowest ask among all exchanges
func BiggestSpread(writer http.ResponseWriter, _ *http.Request) {
	log.Println("Requesting the biggest spread")
	ch := make(chan tickermodels.LimitedJson)

	exchanges.FetchAllExchanges(ch)

	var buyExchange, sellExchange string
	buyPrice := math.MaxFloat64
	sellPrice := 0.0
	for i := 0; i < exchanges.NUMEXCHANGES; i++ {
		exchangeData := <-ch
		for key, val := range exchangeData {
			v, _ := strconv.ParseFloat(val["Ask"], 64)
			if v < buyPrice {
				buyExchange = key
				buyPrice = v
			}

			v, _ = strconv.ParseFloat(val["Bid"], 64)
			if v > sellPrice {
				sellExchange = key
				sellPrice = v
			}
		}
	}

	spread := sellPrice - buyPrice
	response := map[string]string{
		"BuyExchange":  buyExchange,
		"SellExchange": sellExchange,
		"BuyPrice":     strconv.FormatFloat(buyPrice, 'f', -1, 64),
		"SellPrice":    strconv.FormatFloat(sellPrice, 'f', -1, 64),
		"Spread":       strconv.FormatFloat(spread, 'f', -1, 64),
	}
	respond(writer, response, nil)
}

// Return the Bid/Ask data for all exchanges
func AllBidAskData(writer http.ResponseWriter, _ *http.Request) {
	log.Println("Requesting all Bid/Ask data")
	ch := make(chan tickermodels.LimitedJson)

	exchanges.FetchAllExchanges(ch)
	response := make(tickermodels.LimitedJson)

	for i := 0; i < exchanges.NUMEXCHANGES; i++ {
		exchangeData := <-ch
		for key, value := range exchangeData {
			response[key] = value
		}
	}
	respond(writer, response, nil)
}
