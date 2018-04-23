// Package handlers serves as the glue between the exchanges and routes packages
//
// As of now, the meat of the functionality lies in live.go
package handlers

import (
	"encoding/json"
	"github.com/adamhei/honorsproject/exchanges"
	"log"
	"net/http"
)

// Serve the base index route
func Index(writer http.ResponseWriter, _ *http.Request) {
	respond(writer, "{Welcome to the Bitcoin Arbitrage Detector}", nil)
}

// All-in-one function for writing a response or error to the client
func respond(writer http.ResponseWriter, data interface{}, err *exchanges.MyError) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err != nil {
		log.Println(err.Error())
		http.Error(writer, err.Error(), err.Code)
	} else {
		json.NewEncoder(writer).Encode(data)
	}
}
