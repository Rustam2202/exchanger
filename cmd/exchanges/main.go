package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/config"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/logger"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/server"
)

func main() {
	// Interrupt context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// config from confog.yaml
	cfg := config.LoadConfig()
	logger.IntializeLogger(cfg.LoggerConfig)

	httpServer := server.NewServer(cfg.ServerHTTPConfig)
	httpServer.Start(ctx)
}
