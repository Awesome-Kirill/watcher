package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
	"watcher/config"
	"watcher/internal/cache"
	"watcher/internal/site"
	"watcher/internal/status"
	"watcher/internal/transport/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	conf := config.New()

	sites := site.New(conf.PatchFile)

	stat := status.New(conf.Timeout)

	job := cache.New(stat, sites, conf.TTL)
	go job.Watch(ctx)
	e := echo.New()

	h := handler.New(job)
	e.GET("/stat/min", h.GetMin)
	e.GET("/stat/min", h.GetMax)
	e.GET("/stat/:id/site", h.GetSiteStat)

	// Start server
	go func() {
		if err := e.Start(conf.ServerAddress); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	cancel()
	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdown()
	if err := e.Shutdown(ctxShutdown); err != nil {
		e.Logger.Fatal(err)
	}
}
