package routes

import (
	"github.com/adamhei/honorsproject/handlers"
	"github.com/gorilla/mux"
)

// HTTP methods
const (
	GET = "GET"
)

//type route struct {
//	Name, Method, Pattern string
//	HandlerFunc           http.HandlerFunc
//}

//type RouteList []route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods(GET).
		Path("/").
		Name("Index").
		HandlerFunc(handlers.Index)

	router.Methods(GET).
		Path("/live-spread/all").
		Name("AllExchangeData").
		HandlerFunc(handlers.AllBidAskData)

	router.Methods(GET).
		Path("/live-spread/biggest").
		Name("BiggestSpread").
		HandlerFunc(handlers.BiggestSpread)

	router.Methods(GET).
		Path("/live-spread/widest").
		Name("WidestSpread").
		HandlerFunc(handlers.WidestSpread)

	router.Headers("Content-Type", "application/json")
	return router
}
