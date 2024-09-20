package httpserver

import (
	"encoding/json"
	"github.com/dawsonalex/todo-server/build"
	"github.com/dawsonalex/todo-server/config"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
	conf config.C,
) {
	mux.Handle("GET /version", handleVersionGet())
	mux.Handle("GET /config", handleConfigGet(conf))
}

func handleVersionGet() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			e := json.NewEncoder(w)
			err := e.Encode(build.Info())
			if err != nil {
				w.WriteHeader(500)
				_, err = w.Write([]byte(err.Error()))
				if err != nil {
					panic(err)
				}
			}
		},
	)
}

func handleConfigGet(conf config.C) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			e := json.NewEncoder(w)
			err := e.Encode(conf)
			if err != nil {
				w.WriteHeader(500)
				_, err = w.Write([]byte(err.Error()))
				if err != nil {
					panic(err)
				}
			}
		},
	)
}
