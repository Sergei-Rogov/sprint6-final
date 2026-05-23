package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	httpServer *http.Server
	log        *log.Logger
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HTMLHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		httpServer: srv,
		log:        logger,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
