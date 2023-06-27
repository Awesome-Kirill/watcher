package sorted

import "watcher/internal/dto"

type Sort func()

func (c *Sort) MinMax(data map[string]dto.Info) (min, max dto.InfoWithName) {
	siteInfo := make([]dto.InfoWithName, 0, len(data))
	for name, info := range data {
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
	max = siteInfo[0]
	min = siteInfo[0]
	for _, value := range siteInfo {
		if max.ResponseTime < value.ResponseTime {
			max = value
		}
		if min.ResponseTime > value.ResponseTime {
			min = value
		}
	}

	return min, max
}
