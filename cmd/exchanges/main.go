package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/config"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/server"
)

func main() {
	logrus.Info("hello world!")
	
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	wg := &sync.WaitGroup{}

	cfg := config.LoadConfig()
	// logger.IntializeLogger(cfg.LoggerConfig)
	

	httpServer := server.NewServer(cfg.ServerHTTPConfig)
	wg.Add(1)
	go httpServer.Start(ctx, wg)
	wg.Wait()

}
