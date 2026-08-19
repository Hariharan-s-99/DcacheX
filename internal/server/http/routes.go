package http

import "net/http"

func registerRoutes(mux *http.ServeMux, handler *Handler) {
	mux.HandleFunc("/health", handler.Health)
}
	