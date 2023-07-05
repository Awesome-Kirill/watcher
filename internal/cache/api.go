package cache

import (
	"watcher/internal/dto"
)

func (c *Cache) GetURL(url string) (dto.Info, error) {
	c.mu.RLock()
	val, ok := c.data[url]
	c.mu.RUnlock()
	if ok {
		return val, nil
	}

	return dto.Info{}, ErrSiteNotFound
}

func (c *Cache) GetMax() dto.InfoWithName {
	c.sortMux.RLock()
	defer c.sortMux.RUnlock()
	return c.max
}

func (c *Cache) GetMin() dto.InfoWithName {
	c.sortMux.RLock()
	defer c.sortMux.RUnlock()
	return c.min
}
