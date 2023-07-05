package alive

import (
	"bytes"
	"context"
	"net/http"
	"time"
	"watcher/pkg"

	"github.com/rs/zerolog"
)

type Status struct {
	logger *zerolog.Logger
	client *http.Client
}

const durationFalseRequest = -1

func New(timeout time.Duration, logger *zerolog.Logger) *Status {
	return &Status{
		client: &http.Client{
			Timeout: timeout,
		},
		logger: logger,
	}
}
func (c *Status) Alive(ctx context.Context, url string) (bool, time.Duration) {
	start := time.Now()
	c.logger.Debug().Str("url", url).Msg("start")

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, pkg.AddShema(url), bytes.NewReader([]byte{}))
	if err != nil {
		return false, durationFalseRequest
	}
	res, err := c.client.Do(req)
	c.logger.Debug().Str("url", url).Msg("start")
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
