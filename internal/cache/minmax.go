package cache

import (
	"log"
	"watcher/internal/dto"
)

func (c *Cache) minMax() {
	siteInfo := make([]dto.InfoWithName, 0, len(c.data))
	for name, info := range c.data {
		if !info.IsAlive {
			continue
		}
		siteInfo = append(siteInfo, dto.InfoWithName{
			Info: dto.Info{
				IsAlive:      info.IsAlive,
				ResponseTime: info.ResponseTime,
			},
			Name: name,
		})
	}

	if len(siteInfo) == 0 {
		return
	}
	var max = siteInfo[0]
	var min = siteInfo[0]
	for _, value := range siteInfo {
		if max.ResponseTime < value.ResponseTime {
			max = value
		}
		if min.ResponseTime > value.ResponseTime {
			min = value
		}
	}
	c.min = min
	c.max = max
	log.Print("sort finish")
}

func (c *Cache) GetMax() dto.InfoWithName {
	return c.max
}

func (c *Cache) GetMin() dto.InfoWithName {
	return c.min
}
