package status

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

// todo почитай как делать http
// todo http.Client or *http.Client
type Status struct {
	client http.Client
}

func New(timeout time.Duration) *Status {
	return &Status{client: http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},

		Timeout: timeout,
	}}
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
