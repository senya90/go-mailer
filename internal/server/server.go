package server

import (
	"fmt"
	"log/slog"
	"mail-service/internal/handler"
	"net/http"
	"os"
)

type Server struct {
	port    string
	handler *handler.MailHandler
	logger  *slog.Logger
}

func NewServer(port string, handler *handler.MailHandler, logger *slog.Logger) *Server {
	return &Server{
		port:    port,
		handler: handler,
		logger:  logger,
	}
}

func (server *Server) Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /send", server.handler.Send)

	address := fmt.Sprintf(":%s", server.port)
	server.logger.Info("Mail service started on", "port", address)

	err := http.ListenAndServe(address, mux)
	if err != nil {
		server.logger.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
