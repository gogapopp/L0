package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gogapopp/L0/internal/config"
	"github.com/gogapopp/L0/internal/handler"
	cacherecoverer "github.com/gogapopp/L0/internal/libs/cache_recoverer"
	"github.com/gogapopp/L0/internal/logger"
	"github.com/gogapopp/L0/internal/repository/cache"
	"github.com/gogapopp/L0/internal/repository/postgres"
	"github.com/gogapopp/L0/internal/service"
)

func main() {
	var (
		logger   = must(logger.New())
		config   = must(config.New())
		postgres = must(postgres.New(config))
		cache    = cache.New(time.Hour*24, time.Hour*24)
		service  = service.New(postgres, cache)
	)
	defer postgres.Close()

	if err := postgres.MigrateUp(config); err != nil {
		logger.Fatal(err)
	}

	cacherecoverer.CacheRecover(logger, cache, postgres)

	mux := http.DefaultServeMux
	mux.HandleFunc("GET /orders/{id}", handler.GetOrderById(logger, service))

	httpserver := &http.Server{
		Addr:         config.Address,
		Handler:      mux,
		IdleTimeout:  config.IdleTimeout,
		ReadTimeout:  config.Timeout,
		WriteTimeout: config.Timeout,
	}

	if err := httpserver.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-quit

	if err := postgres.MigrateDown(config); err != nil {
		logger.Fatal(err)
	}

	if err := httpserver.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
		logger.Fatal(err)
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}