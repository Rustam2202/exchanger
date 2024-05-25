package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rustam2202/exchanger/internal/config"
	"github.com/Rustam2202/exchanger/internal/logger"
	"github.com/Rustam2202/exchanger/internal/server"
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
