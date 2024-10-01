package router

import (
	"lavajato/src/router/routes"

	"github.com/gorilla/mux"
)

func ToGenerate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
