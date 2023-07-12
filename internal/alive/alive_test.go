//go:build tests

package alive

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestStatus_Alive(t *testing.T) {
	type fields struct {
		client *http.Client
		logger *zerolog.Logger
	}
	type args struct {
		ctx context.Context
		url string
	}
	var s io.Writer
	l := zerolog.New(s)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  time.Duration
	}{
		{
			name: "error", fields: fields{
				client: http.DefaultClient,
				logger: &l,
			}, args: args{
				ctx: context.Background(),
				url: "https://error.error.1232dfsdfdsfsdf.com",
			},
			want:  false,
			want1: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Status{
				logger: tt.fields.logger,
				client: tt.fields.client,
			}
			got, got1 := c.Alive(tt.args.ctx, tt.args.url)
			assert.Equal(t, got, tt.want, "Alive() they should be equal")
			assert.Equal(t, got1, tt.want1, "Alive() they should be equal")
		})
	}
}
