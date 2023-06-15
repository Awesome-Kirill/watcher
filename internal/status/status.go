package status

import (
	"context"
	"log"
	"net/http"
	"time"
)

// todo почитай как делать http Request
type Status struct {
	client *http.Client
}

func New(client *http.Client) *Status {
	return &Status{client: client}
}
func (c *Status) GetStatus(ctx context.Context, url string) (bool, time.Duration) {
	start := time.Now()
	log.Printf("reqest start:%v", url)
	res, err := c.client.Head("https://" + url)
	log.Printf("reqest finish:%v", url)
	finish := time.Since(start)
	if err != nil {
		return false, finish
	}

	defer res.Body.Close()

	if res.StatusCode < 300 {
		return true, finish
	}

	return false, finish
}
