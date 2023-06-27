package cache

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
	"watcher/internal/dto"
)

type Aliver interface {
	Alive(context.Context, string) (isAlive bool, responseTime time.Duration)
}

// todo видео про именование интерфейса
type Hoster interface {
	Hosts(context.Context) ([]string, error)
}

type Cache struct {
	ttl    time.Duration
	aliver Aliver
	hoster Hoster

	mu   sync.Mutex
	data map[string]dto.Info

	min, max dto.InfoWithName
}

func (c *Cache) update(ctx context.Context) {
	log.Print("ticker start")
	var wg sync.WaitGroup
	sits, err := c.hoster.Hosts(ctx)

	// todo
	if err != nil {
		log.Print("Hosts err")
		return
	}
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
}

// todo error when
// todo start ticker
func (c *Cache) Watch(ctx context.Context) {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for {
		c.update(ctx)
		c.minMax()
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			return
		case <-ticker.C:
			continue
		}
	}
}

var SiteNotFound = errors.New("SiteNotFound")

func (c *Cache) GetUrl(url string) (dto.Info, error) {
	val, ok := c.data[url]
	if ok {
		return val, nil
	}

	return dto.Info{}, SiteNotFound
}

func New(stater Aliver, sited Hoster, ttl time.Duration) *Cache {
	return &Cache{
		ttl:    ttl,
		aliver: stater,
		hoster: sited,
		mu:     sync.Mutex{},
		data:   make(map[string]dto.Info),
	}
}
