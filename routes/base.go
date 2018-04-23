// Package routes provides a mux Router with a directory of all supported endpoints and their corresponding handler function
package routes

import (
	"github.com/adamhei/honorsproject/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods(http.MethodGet).
		Path("/").
		Name("Index").
		HandlerFunc(handlers.Index)

	router.Methods(http.MethodGet).
		Path("/live-spread/all").
		Name("AllExchangeData").
		HandlerFunc(handlers.AllBidAskData)

	router.Methods(http.MethodGet).
		Path("/live-spread/widest").
		Name("WidestSpread").
		HandlerFunc(handlers.WidestSpread)

	router.Headers("Content-Type", "application/json")
	return router
}
