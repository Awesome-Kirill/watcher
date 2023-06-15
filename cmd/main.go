package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/signal"
	"watcher/config"
	"watcher/internal/cache"
	"watcher/internal/site"
	"watcher/internal/status"
	"watcher/internal/transport/handler"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	conf := config.New()

	sites := site.New(conf.PatchFile)

	c := http.DefaultClient
	c.Timeout = conf.Timeout
	stat := status.New(c)

	job := cache.New(stat, sites, conf.TTL)
	go job.Watch(ctx)
	e := echo.New()

	h := handler.New(job)
	e.GET("/stat/min", h.GetMin)
	e.GET("/stat/min", h.GetMax)
	e.GET("/stat/:id/site", h.GetSiteStat)

	e.Start(":1323")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	s := <-ch
	cancel()
	fmt.Println("Got signal:", s)
}
