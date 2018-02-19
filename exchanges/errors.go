package exchanges

import (
	"github.com/adamhei/honorsproject/exchanges/tickermodels"
	"log"
)

type MyError struct {
	Err  string
	Code int
}

func (e *MyError) Error() string {
	return e.Err
}

func (e *MyError) ErrorCode() int {
	return e.Code
}

// Print the error message and send an empty response through the channel
func errorHandler(errorMsg string, ch chan<- tickermodels.LimitedJson) {
	log.Print(errorMsg)
	ch <- make(tickermodels.LimitedJson)
}
