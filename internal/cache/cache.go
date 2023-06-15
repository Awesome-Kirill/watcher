package cache

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

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
	stater Stater
	sited  Sited

	mu   sync.Mutex
	data map[string]Info

	min, max InfoWithName
}

type Stater interface {
	GetStatus(context.Context, string) (isAlive bool, responseTime time.Duration)
}

type Sited interface {
	GetSites(context.Context) ([]string, error)
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
			isAlive, finish := c.stater.GetStatus(ctx, url)
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

func (c *Cache) GetMax() InfoWithName {
	return c.max
}

func (c *Cache) GetMin() InfoWithName {
	return c.min
}

var SiteNotFound = errors.New("SiteNotFound")

func (c *Cache) GetUrl(url string) (Info, error) {
	val, ok := c.data[url]
	if ok {
		return val, nil
	}

	return Info{}, SiteNotFound
}
func (c *Cache) minMax() {
	siteInfo := make([]InfoWithName, 0, len(c.data))
	for name, info := range c.data {
		siteInfo = append(siteInfo, InfoWithName{
			Info: Info{
				IsAlive:      info.IsAlive,
				ResponseTime: info.ResponseTime,
			},
			Name: name,
		})
	}

	var max = siteInfo[0]
	var min = siteInfo[0]
	for _, value := range siteInfo {
		if max.ResponseTime < value.ResponseTime && value.IsAlive {
			max = value
		}
		if min.ResponseTime > value.ResponseTime && value.IsAlive {
			min = value
		}
	}
	c.min = min
	c.max = max
	log.Print("sort finish")
}

func New(stater Stater, sited Sited, ttl time.Duration) *Cache {
	return &Cache{
		ttl:    ttl,
		stater: stater,
		sited:  sited,
		mu:     sync.Mutex{},
		data:   make(map[string]Info),
	}
}
