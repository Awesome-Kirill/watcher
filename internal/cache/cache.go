package cache

import (
	"context"
	"errors"
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
	Hosts(context.Context) ([]string, error)
}

type sorter interface {
	MinMax(map[string]dto.Info) (min, max dto.InfoWithName)
}
type Cache struct {
	aliver aliver
	hoster hoster
	sorter sorter

	mu   *sync.Mutex
	data map[string]dto.Info

	min, max dto.InfoWithName
	ttl      time.Duration
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
		// c.minMax()
		c.min, c.max = c.sorter.MinMax(c.data)
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
		ttl:    ttl,
		aliver: aliver,
		hoster: hoster,
		sorter: sorter,
		mu:     &sync.Mutex{},
		data:   make(map[string]dto.Info),
	}
}
