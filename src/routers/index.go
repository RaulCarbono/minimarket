package routers

import (
	"github.com/go/mini_market/src/server"
	"github.com/gorilla/mux"
)

func BindRoutes(s server.Server, r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use()
	HomeRouter(s, r)
}
