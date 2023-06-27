package alive

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestStatus_Alive(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  time.Duration
	}{
		{
			name: "error", fields: fields{client: http.DefaultClient}, args: args{
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
				client: tt.fields.client,
			}
			got, got1 := c.Alive(tt.args.ctx, tt.args.url)
			if got != tt.want {
				t.Errorf("Alive() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Alive() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
