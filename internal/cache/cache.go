package cache

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

type Aliver interface {
	Alive(context.Context, string) (isAlive bool, responseTime time.Duration)
}

type Sited interface {
	GetSites(context.Context) ([]string, error)
}

type Info struct {
	IsAlive      bool
	ResponseTime time.Duration
}

type InfoWithName struct {
	Info
	Name string
}
type Cache struct {
	ttl    time.Duration
	aliver Aliver
	sited  Sited

	mu   sync.Mutex
	data map[string]Info

	min, max InfoWithName
}

func (c *Cache) update(ctx context.Context) {
	log.Print("ticker start")
	var wg sync.WaitGroup
	sits, err := c.sited.GetSites(ctx)

	// todo
	if err != nil {
		log.Print("GetSites err")
		return
	}
	for _, url := range sits {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			isAlive, finish := c.aliver.Alive(ctx, url)
			c.mu.Lock()
			c.data[url] = Info{
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

func (c *Cache) GetUrl(url string) (Info, error) {
	val, ok := c.data[url]
	if ok {
		return val, nil
	}

	return Info{}, SiteNotFound
}

func New(stater Aliver, sited Sited, ttl time.Duration) *Cache {
	return &Cache{
		ttl:    ttl,
		aliver: stater,
		sited:  sited,
		mu:     sync.Mutex{},
		data:   make(map[string]Info),
	}
}
