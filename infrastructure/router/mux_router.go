package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

var muxRouterInstance = mux.NewRouter()

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) Serve(port string) {
	err := http.ListenAndServe(port, muxRouterInstance)
	if err != nil {
		return
	}
}
