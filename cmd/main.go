package main

import (
	"mail-service/internal/config"
	"mail-service/internal/handler"
	"mail-service/internal/mailer"
	"mail-service/internal/server"
)

func main() {
	config := config.LoadEnv()

	mailer := mailer.NewMailer(config)
	handler := handler.NewMailHandler(mailer)
	server := server.NewServer(config.Port, handler)

	server.Run()
}
