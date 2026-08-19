package http

import (
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer(addr string) *Server {
	handler := NewHandler()

	mux := http.NewServeMux()

	registerRoutes(mux, handler)

	handlerWithMiddleware := Logging(mux)

	return &Server{
		server: &http.Server{
			Addr:         addr,
			Handler:      handlerWithMiddleware,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
