package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"watcher/config"
	_ "watcher/docs"
	"watcher/internal/alive"
	cacheStatus "watcher/internal/cache"
	"watcher/internal/file"
	"watcher/internal/sorted"
	"watcher/internal/transport/handler"

	"github.com/rs/zerolog"
)

// @title       watcher
// @version     1.0
// @description Сервис для проверки доступности сайтов
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	conf := config.New()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if conf.DebugMod {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	sites := file.Load(conf.PatchFile, &log)
	stat := alive.New(conf.Timeout, &log)
	cache := cacheStatus.New(new(sorted.Sort), stat, sites, &log, conf.TTL)

	go func() {
		cache.Watch(ctx)
	}()

	h := handler.New(cache, conf.AdminKey)

	// Start server
	go func() {
		if err := h.Start(conf.ServerAddress); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("shutting down the server")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	cancel()

	//nolint:gomnd // explanation
	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	if err := h.Shutdown(ctxShutdown); err != nil {
		//nolint:gocritic // explanation
		log.Fatal().Err(err).Msg("server stop error")
	}
	log.Info().Msg("stopped!")
}
