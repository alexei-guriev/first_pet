package httpapi

import (
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

func New(handler http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:              ":8080",
			Handler:           handler,
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       30 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       60 * time.Second,
		},
	}
}

func (s *Server) Run() {
	s.srv.ListenAndServe()
}
