package alive

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"time"
	"watcher/pkg"
)

type Status struct {
	client *http.Client
}

const durationFalseRequest = -1

func New(timeout time.Duration) *Status {
	return &Status{client: &http.Client{
		Timeout: timeout,
	}}
}
func (c *Status) Alive(ctx context.Context, url string) (bool, time.Duration) {
	start := time.Now()
	log.Printf("reqest start:%v", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, pkg.AddShema(url), bytes.NewReader([]byte{}))
	if err != nil {
		return false, durationFalseRequest
	}
	res, err := c.client.Do(req)
	log.Printf("reqest finish:%v", url)
	finish := time.Since(start)
	if err != nil {
		return false, durationFalseRequest
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return true, finish
	}

	return false, durationFalseRequest
}
