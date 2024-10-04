package routers

import (
	"net/http"

	"github.com/go/mini_market/src/handlers"
	"github.com/go/mini_market/src/server"
	"github.com/gorilla/mux"
)

func HomeRouter(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
}
