package handlers

import (
	"github.com/adamhei/honorsproject/exchanges"
	"github.com/adamhei/honorsproject/tickermodels"
	"log"
	"math"
	"net/http"
	"strconv"
)

// WidestSpread pulls all exchange ticker data, calculate the largest price disparity between two exchanges, and
// writes back all information necessary for the client to display the opportunity
//
// Note: the excessive string parsing is a necessary evil of how we return data to the client
func WidestSpread(writer http.ResponseWriter, _ *http.Request) {
	log.Println("Requesting the widest spread")

	ch := make(chan tickermodels.Ticker)
	exchanges.FetchAllExchanges(ch)

	var buyTicker, sellTicker tickermodels.Ticker
	bestEffectiveBuy := math.MaxFloat64
	bestEffectiveSell := 0.0

	for i := 0; i < exchanges.NumExchanges; i++ {
		ticker := <-ch
		if ticker.Exchange == "" {
			// Skip over exchanges whose data we could not fetch
			continue
		}

		ask, _ := strconv.ParseFloat(ticker.Ask, 64)
		askSize, _ := strconv.ParseFloat(ticker.AskSize, 64)
		taker, _ := strconv.ParseFloat(ticker.TakerFee, 64)
		taker /= 100
		withdrawal, _ := strconv.ParseFloat(ticker.BTCWithdrawalFee, 64)

		// Save the buy exchange if it has the new lowest effective price
		effectiveBuyPrice := (ask * askSize) / (askSize*(1-taker) - withdrawal)
		if effectiveBuyPrice < bestEffectiveBuy {
			bestEffectiveBuy = effectiveBuyPrice
			buyTicker = ticker
		}

		// Save the sell exchange if it has the new highest effective price
		bid, _ := strconv.ParseFloat(ticker.Bid, 64)
		effectiveSellPrice := bid * (1 - taker)
		if effectiveSellPrice > bestEffectiveSell {
			bestEffectiveSell = effectiveSellPrice
			sellTicker = ticker
		}
	}

	bestAsk, _ := strconv.ParseFloat(buyTicker.Ask, 64)
	bestAskQty, _ := strconv.ParseFloat(buyTicker.AskSize, 64)
	bestBid, _ := strconv.ParseFloat(sellTicker.Bid, 64)
	bestBidQty, _ := strconv.ParseFloat(sellTicker.BidSize, 64)

	buyWithdrawalFee, _ := strconv.ParseFloat(buyTicker.BTCWithdrawalFee, 64)
	buyTakerFee, _ := strconv.ParseFloat(buyTicker.TakerFee, 64)
	buyTakerFee /= 100
	sellTakerFee, _ := strconv.ParseFloat(sellTicker.TakerFee, 64)
	sellTakerFee /= 100

	// Find the theoretical profit
	var profit float64
	effectivePurchaseQty := bestAskQty*(1-buyTakerFee) - buyWithdrawalFee
	if effectivePurchaseQty <= bestBidQty {
		// We can buy less than the bid size on the sell exchange
		invested := bestAsk * bestAskQty
		returned := effectivePurchaseQty * bestBid * (1 - sellTakerFee)
		profit = returned - invested
	} else {
		// We can buy more than the bid size on the sell exchange
		invested := bestAsk * bestBidQty
		effectivePurchaseQty = bestBidQty*(1-buyTakerFee) - buyWithdrawalFee
		returned := effectivePurchaseQty * bestBid * (1 - sellTakerFee)
		profit = returned - invested
	}

	// Build the response
	btcQuantity := math.Min(bestAskQty, bestBidQty)
	response := map[string]interface{}{
		"buyExchange":      buyTicker.Exchange,
		"buyPrice":         buyTicker.Ask,
		"buyQuantity":      buyTicker.AskSize,
		"buyTakerFee":      buyTicker.TakerFee,
		"buyWithdrawalFee": buyTicker.BTCWithdrawalFee,
		"hasWithdrawalFee": buyTicker.HasWithdrawalFee,

		"sellExchange": sellTicker.Exchange,
		"sellPrice":    sellTicker.Bid,
		"sellQuantity": sellTicker.BidSize,
		"sellTakerFee": sellTicker.TakerFee,

		"btcQuantity": strconv.FormatFloat(btcQuantity, 'f', -1, 64),
		"profit":      strconv.FormatFloat(profit, 'f', 2, 64),
	}
	respond(writer, response, nil)
}

// Return the Bid/Ask data for all exchanges
func AllBidAskData(writer http.ResponseWriter, _ *http.Request) {
	log.Println("Requesting all Bid/Ask data")
	ch := make(chan tickermodels.Ticker)

	exchanges.FetchAllExchanges(ch)
	response := make([]tickermodels.Ticker, 0)

	for i := 0; i < exchanges.NumExchanges; i++ {
		ticker := <-ch
		if ticker.Exchange != "" {
			// There is actual ticker data
			response = append(response, ticker)
		}
	}

	respond(writer, response, nil)
}
