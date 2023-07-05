package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
	"watcher/config"
	"watcher/internal/alive"
	cacheStatus "watcher/internal/cache"
	"watcher/internal/file"
	"watcher/internal/sorted"
	"watcher/internal/transport/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	conf := config.New()

	sites := file.Load(conf.PatchFile)
	stat := alive.New(conf.Timeout)
	cache := cacheStatus.New(new(sorted.Sort), stat, sites, conf.TTL)

	go cache.Watch(ctx)

	e := echo.New()
	h := handler.New(cache)

	e.GET("/stat/min", h.GetMin)
	e.GET("/stat/max", h.GetMax)
	e.GET("/stat/:id/site", h.GetSiteStat)

	e.GET("/admin/stat", h.GetStat, middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == conf.AdminKey, nil
	}))
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
	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), conf.TimeoutShutdown*time.Second)
	defer cancelShutdown()
	if err := e.Shutdown(ctxShutdown); err != nil {
		e.Logger.Fatal(err)
	}
}
