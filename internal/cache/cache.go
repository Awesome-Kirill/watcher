package cache

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
	"watcher/internal/dto"
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
}

func (c *Cache) update(ctx context.Context) error {
	log.Print("ticker start")
	sits, err := c.hoster.Host(ctx)
	// todo
	if err != nil {
		log.Print("Host err")
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
	log.Print("ticker finish")
	return nil
}

func (c *Cache) Watch(ctx context.Context) {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for {
		err := c.update(ctx)
		if err != nil {
			log.Printf("err update:%v", ctx.Err())
			continue
		}

		c.sortMux.Lock()
		c.min, c.max = c.sorter.MinMax(c.data)
		c.sortMux.Unlock()
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			return
		case <-ticker.C:
			continue
		}
	}
}

func New(sorter sorter, aliver aliver, hoster hoster, ttl time.Duration) *Cache {
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
	}
}
