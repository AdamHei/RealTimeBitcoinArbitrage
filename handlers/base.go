package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/adamhei/honorsproject/exchanges"
)

func Index(writer http.ResponseWriter, _ *http.Request) {
	respond(writer, "{Welcome to the Bitcoin Arbitrage Detector}", nil)
}

func respond(writer http.ResponseWriter, data interface{}, err *exchanges.MyError) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err != nil {
		log.Println(err.Error())
		http.Error(writer, err.Error(), err.Code)
	} else {
		json.NewEncoder(writer).Encode(data)
	}
}
