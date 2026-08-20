package http

import "net/http"

func registerRoutes(mux *http.ServeMux, handler *Handler) {
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /cache/{key}", handler.Get)
	mux.HandleFunc("PUT /cache{key}", handler.Get)
	mux.HandleFunc("DELETE /cache{key}", handler.Delete)
}

