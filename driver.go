package main

import (
	"github.com/adamhei/honorsproject/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":81", router))
}
