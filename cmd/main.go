package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"

	"github.com/MaximTretjakov/nofelet-web/config"
	"github.com/MaximTretjakov/nofelet-web/internal/app/web"
	"github.com/MaximTretjakov/nofelet-web/internal/dependency"
	"github.com/MaximTretjakov/nofelet-web/internal/storage/postgres"
	"github.com/MaximTretjakov/nofelet-web/pkg/httpserver"
)

func main() {
	// Создаем конфиг
	if err := config.New(); err != nil {
		panic(err)
	}
	cfg := config.Current()

	// Создаем логгер
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Задаем режим
	gin.SetMode(gin.ReleaseMode)
	if cfg.Debug {
		logger.Info("gin debug mode enabled")
		gin.SetMode(gin.DebugMode)
	}

	// Создаем современный экспортер v0.66.0
	exporter, err := otelprom.New()
	if err != nil {
		log.Fatalf("failed to create exporter: %v", err)
	}

	provider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(exporter))
	defer func() { _ = provider.Shutdown(context.Background()) }()

	// Создаем коннекшен к бд
	db, dbErr := postgres.New(cfg.DB.ConnectionString)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer func() {
		if dbErr = db.Close(); dbErr != nil {
			logger.Error("error:", slog.Any("db init error:", dbErr))
		}
	}()

	// Создаем DI
	deps, depErr := dependency.New(&cfg, logger, db)
	if depErr != nil {
		log.Fatal(depErr)
	}

	// Создаем контейнер приложения
	if sigErr := web.New(deps); sigErr != nil {
		log.Fatal(sigErr)
	}

	// Создаем и запускаем сервер
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	httpServer := httpserver.New(deps.Web.Routes,
		httpserver.WithAddress(cfg.Web.Port),
		httpserver.WithServerCRT(cfg.Crt),
		httpserver.WithServerKey(cfg.Key),
		httpserver.WithReadTimeout(cfg.Web.ReadTimeout),
		httpserver.WithReadHeaderTimeout(cfg.Web.ReadHeaderTimeout),
		httpserver.WithWriteTimeout(cfg.Web.WriteTimeout),
		httpserver.WithShutdownTimeout(cfg.Web.ShutdownTimeout),
	)

	// Graceful shutdown
	select {
	case s := <-interrupt:
		logger.Error("error", slog.String("signal", s.String()))
	case err := <-httpServer.Notify():
		logger.Error("httpServer.Notify", slog.Any("error", err))
	}

	if err := httpServer.Shutdown(); err != nil {
		logger.Error("httpServer.Shutdown", slog.Any("error", err))
	}
}
