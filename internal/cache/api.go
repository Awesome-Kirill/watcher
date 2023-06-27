package cache

import (
	"watcher/internal/dto"
)

func (c *Cache) GetUrl(url string) (dto.Info, error) {
	val, ok := c.data[url]
	if ok {
		return val, nil
	}

	return dto.Info{}, SiteNotFound
}

func (c *Cache) GetMax() dto.InfoWithName {
	return c.max
}

func (c *Cache) GetMin() dto.InfoWithName {
	return c.min
}
