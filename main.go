package main

import (
	"github/chrislim1914/go-xumm/api"
	"log"
	"net/http"
)

func main() {
	router := api.NewRouter(api.AllRoutes())
	router.GlobalOPTIONS = api.CorsMiddleware()
	log.Fatal(http.ListenAndServe(":8080", router))
}
