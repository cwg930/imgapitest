package routers

import (
	"net/http"
	"github.com/gorilla/mux"
//	"github.com/cwg930/imgapitest/controllers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		
		handler = route.HandlerFunc
	//	handler = controllers.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
