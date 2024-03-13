package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,          
		ReadTimeout:    20 * time.Second, 
		WriteTimeout:   20 * time.Second, 
	}

	return s.httpServer.ListenAndServe()
}