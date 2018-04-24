package exchanges

import (
	"github.com/adamhei/honorsproject/tickermodels"
	"log"
)

// Custom error types to include http response codes

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
func errorHandler(customErr string, err error, ch chan<- tickermodels.Ticker) {
	log.Println(customErr)
	log.Println(err)
	ch <- *new(tickermodels.Ticker)
}
