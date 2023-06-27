package alive

import (
	"bytes"
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
func (c *Status) Alive(ctx context.Context, url string) (bool, time.Duration) {
	start := time.Now()
	log.Printf("reqest start:%v", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, bytes.NewReader([]byte{}))
	if err != nil {
		return false, -1
	}
	res, err := c.client.Do(req)
	//res, err := c.client.Head("https://" + url)
	log.Printf("reqest finish:%v", url)
	finish := time.Since(start)
	if err != nil {
		return false, -1
	}

	defer res.Body.Close()

	if res.StatusCode < 300 {
		return true, finish
	}

	return false, -1
}
