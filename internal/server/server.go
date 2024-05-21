package server

import (
	"fmt"
	"net/http"
)

// Server представляє сервер з використанням бібліотеки Chi.
type Server struct {
	server *http.Server
}

func NewServer(handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	fmt.Printf("Server run :%v", s.server.Addr)
	return http.ListenAndServe(s.server.Addr, s.server.Handler)
}
