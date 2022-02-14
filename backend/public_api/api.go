package public_api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Api struct {
	Router *mux.Router
}

func Init() error {
	api := Api{}
	api.Router = mux.NewRouter()

	api.RegisterRoutes()
	return api.Serve()
}

func (api *Api) RegisterRoutes() {
	api.Router.HandleFunc("/hello", HelloWorldHandler)
}

func (api *Api) Serve() error {
	server := http.Server{
		Addr:    ":8080",
		Handler: api.Router,
	}
	return server.ListenAndServe()
}
