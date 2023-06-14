package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	IsAlive      bool
	ResponseTime time.Duration
}

func GetSiteList() []string {
	return []string{
		"google.com", "youtube.com", "facebook.com", "baidu.com", "wikipedia.org", "qq.com", "taobao.com", "yahoo.com",
		"tmall.com", "amazon.com", "google.co.in", "twitter.com", "sohu.com", "jd.com", "live.com", "instagram.com",
		"sina.com.cn", "weibo.com", "google.co.jp", "reddit.com", "vk.com", "360.cn", "login.tmall.com", "blogspot.com",
		"yandex.ru", "google.com.hk", "netflix.com", "linkedin.com", "pornhub.com", "google.com.br", "twitch.tv",
		"pages.tmall.com", "csdn.net", "yahoo.co.jp", "mail.ru", "aliexpress.com", "alipay.com", "office.com",
		"google.fr", "google.ru", "google.co.uk", "microsoftonline.com", "google.de", "ebay.com", "microsoft.com",
		"livejasmin.com", "t.co", "bing.com", "xvideos.com", "google.ca",
	}
}
func main() {

	// todo Pointer or Value
	cache := make(map[string]*Result)

	for _, s := range GetSiteList() {

		cache[s] = &Result{}
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {

		select {
		case <-ticker.C:
			wg := sync.WaitGroup{}
			client := http.DefaultClient
			client.Timeout = 1 * time.Second
			var mu sync.Mutex
			for _, key := range GetSiteList() {
				wg.Add(1)
				go func(key string) {
					defer wg.Done()
					start := time.Now()

					res, err := client.Head("https://" + key)

					total := time.Since(start)
					if err != nil {
						return
					}
					defer res.Body.Close()

					var isAlive bool
					if res.StatusCode <= 300 {
						isAlive = true
					}

					mu.Lock()
					cache[key] = &Result{
						IsAlive:      isAlive,
						ResponseTime: total,
					}
					mu.Unlock()
				}(key)

			}

			wg.Wait()
			fmt.Printf("Done")
		}
	}
}
