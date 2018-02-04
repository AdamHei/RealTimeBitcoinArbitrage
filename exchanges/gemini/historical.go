package gemini

import (
	"encoding/json"
	"fmt"
	"github.com/adamhei/historicaldata/trademodels"
	"net/http"
	"time"
)

const historyUrl = "https://api.gemini.com/v1/trades/btcusd?since=%d&limit_trades=500"

// Deprecated
func GetTradeHistory(from time.Time, to time.Time) []trademodels.GeminiOrder {
	orders := make([]trademodels.GeminiOrder, 0)
	for indexTime := from; indexTime.Before(to); {
		formattedUrl := fmt.Sprintf(historyUrl, indexTime.Unix())

		resp, err := http.Get(formattedUrl)
		if err != nil {
			fmt.Println(err)
			return make([]trademodels.GeminiOrder, 0)
		}

		nextOrders := make([]trademodels.GeminiOrder, 0)
		err = json.NewDecoder(resp.Body).Decode(&nextOrders)
		resp.Body.Close()

		if err != nil {
			fmt.Println(err)
			return make([]trademodels.GeminiOrder, 0)
		}

		orders = append(nextOrders, orders...)
		indexTime = time.Unix(0, orders[0].TimestampMs*int64(time.Millisecond))
	}

	return orders
}
