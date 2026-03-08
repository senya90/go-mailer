package main

import (
	"mail-service/internal/config"
	"mail-service/internal/handler"
	"mail-service/internal/logger"
	"mail-service/internal/mailer"
	"mail-service/internal/server"
)

func main() {
	config := config.LoadEnv()
	logger := logger.NewLogger(config.IsProduction)

	mailer := mailer.NewMailer(config, logger)
	handler := handler.NewMailHandler(mailer, logger)
	server := server.NewServer(config.Port, handler, logger)

	server.Run()
}
