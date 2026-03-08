package server

import (
	"fmt"
	"log"
	"mail-service/internal/handler"
	"net/http"
)

type Server struct {
	port    string
	handler *handler.MailHandler
}

func NewServer(port string, handler *handler.MailHandler) *Server {
	return &Server{
		port:    port,
		handler: handler,
	}
}

func (server *Server) Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/send", server.handler.Send)

	address := fmt.Sprintf(":%s", server.port)
	log.Printf("Mail service started on %s", address)

	err := http.ListenAndServe(address, mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
