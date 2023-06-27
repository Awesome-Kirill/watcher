package alive

import (
	"bytes"
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"time"
	"watcher/pkg"
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

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, pkg.AddShema(url), bytes.NewReader([]byte{}))
	if err != nil {
		return false, -1
	}
	res, err := c.client.Do(req)
	log.Printf("reqest finish:%v", url)
	finish := time.Since(start)
	if err != nil {
		return false, -1
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return true, finish
	}

	return false, -1
}
