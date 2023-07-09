package cache

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	"watcher/internal/dto"

	"github.com/rs/zerolog"
)

// ErrSiteNotFound todo
var ErrSiteNotFound = errors.New("site not found")

type aliver interface {
	Alive(context.Context, string) (isAlive bool, responseTime time.Duration)
}

type hoster interface {
	Host(context.Context) ([]string, error)
}

type sorter interface {
	MinMax(map[string]dto.Info) (min, max dto.InfoWithName)
}
type Cache struct {
	aliver aliver
	hoster hoster
	sorter sorter

	mu   *sync.RWMutex
	data map[string]dto.Info

	sortMux  *sync.RWMutex
	min, max dto.InfoWithName
	ttl      time.Duration

	logger *zerolog.Logger
}

func (c *Cache) update(ctx context.Context) error {
	c.logger.Info().Msg("update start")
	sits, err := c.hoster.Host(ctx)
	// todo
	if err != nil {
		c.logger.Err(err).Msg("err get host")
		return fmt.Errorf("err hoster.Host:%w", err)
	}

	var wg sync.WaitGroup
	for _, url := range sits {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			isAlive, finish := c.aliver.Alive(ctx, url)
			c.mu.Lock()
			c.data[url] = dto.Info{
				IsAlive:      isAlive,
				ResponseTime: finish,
			}
			c.mu.Unlock()
		}(url)
	}

	wg.Wait()
	c.logger.Info().Msg("update finish")
	return nil
}

func (c *Cache) Watch(ctx context.Context) {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for {
		err := c.update(ctx)
		if err != nil {
			c.logger.Err(ctx.Err()).Msg("err update")
			continue
		}

		c.sortMux.Lock()
		c.min, c.max = c.sorter.MinMax(c.data)
		c.sortMux.Unlock()
		select {
		case <-ctx.Done():
			c.logger.Err(ctx.Err()).Msg("watcher context done")
			return
		case <-ticker.C:
			continue
		}
	}
}

func New(sorter sorter, aliver aliver, hoster hoster, logger *zerolog.Logger, ttl time.Duration) *Cache {
	return &Cache{
		aliver:  aliver,
		hoster:  hoster,
		sorter:  sorter,
		mu:      &sync.RWMutex{},
		data:    make(map[string]dto.Info),
		sortMux: &sync.RWMutex{},
		min:     dto.InfoWithName{},
		max:     dto.InfoWithName{},
		ttl:     ttl,
		logger:  logger,
	}
}
