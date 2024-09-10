package httpserver

import "net/http"

func New() http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
	)
	var handler http.Handler = mux
	return handler
}
