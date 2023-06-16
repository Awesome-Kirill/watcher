package cache

import "log"

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

func (c *Cache) GetMax() InfoWithName {
	return c.max
}

func (c *Cache) GetMin() InfoWithName {
	return c.min
}
