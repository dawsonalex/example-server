package httpserver

import (
	"encoding/json"
	"github.com/dawsonalex/todo-server/build"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
) {
	mux.Handle("GET /version", handleVersionGet())
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
