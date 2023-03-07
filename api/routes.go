package api

import "github.com/julienschmidt/httprouter"

type Route struct {
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
		Route{"GET", "/xumm-server", CheckXummServer},
		Route{"POST", "/xumm-signin", SignIn},
		Route{"POST", "/xumm-payment", SendPayment},
	}
	return routes
}
