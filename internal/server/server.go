package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/docs"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/logger"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/server/handlers"
	"go.uber.org/zap"
)

type Server struct {
	cfg             *ServerHTTPConfig
	exchReqHandler  handlers.ExchangeRequest
	exchRespHandler handlers.ExchangeResponse
	HttpServer      *http.Server
}

func NewServer(cfg ServerHTTPConfig) *Server {
	return &Server{
		cfg: &cfg,
	}
}

// @title			Exchanger API
// @description	Exchanger app server.
// @BasePath		/
func (s *Server) Start(ctx context.Context) {
	wg := &sync.WaitGroup{}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world from Exchanger HTTP server")
	})

	router.POST("/exchange", handlers.ExchangePost)

	// swagger API documentation
	router.GET("/docs/index.html", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.HttpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port),
		Handler: router,
	}

	// graceful shutdown
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := s.HttpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Logger.Error("Failed to start HTTP server", zap.Error(err))
		}
	}()
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	cancel()
	if err := s.HttpServer.Shutdown(shutdownCtx); err != nil {
		logger.Logger.Error("Failed to shutdown HTTP server", zap.Error(err))
	}
	logger.Logger.Info("Shutting down HTTP server ...")
	wg.Wait()
}
