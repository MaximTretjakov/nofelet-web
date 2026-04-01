package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/MaximTretjakov/nofelet-web/config"
	"github.com/MaximTretjakov/nofelet-web/internal/app/web"
	"github.com/MaximTretjakov/nofelet-web/internal/dependency"
	"github.com/MaximTretjakov/nofelet-web/pkg/httpserver"
)

func main() {
	if err := config.New(); err != nil {
		panic(err)
	}
	cfg := config.Current()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	deps, depErr := dependency.New(&cfg, logger)
	if depErr != nil {
		log.Fatal(depErr)
	}

	if sigErr := web.New(deps); sigErr != nil {
		log.Fatal(sigErr)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	httpServer := httpserver.New(deps.Web.Routes,
		httpserver.WithAddress(cfg.Web.Port),
		httpserver.WithServerCRT(cfg.Web.ServerCrt),
		httpserver.WithServerKey(cfg.Web.ServerKey),
		httpserver.WithReadTimeout(cfg.Web.ReadTimeout),
		httpserver.WithReadHeaderTimeout(cfg.Web.ReadHeaderTimeout),
		httpserver.WithWriteTimeout(cfg.Web.WriteTimeout),
		httpserver.WithShutdownTimeout(cfg.Web.ShutdownTimeout),
	)

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
