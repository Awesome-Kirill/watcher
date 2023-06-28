package cache

import (
	"context"
	"reflect"
	"sync"
	"testing"
	"time"
	"watcher/internal/dto"
	"watcher/internal/sorted"
)

type MockHoster struct {
}

func (h *MockHoster) Host(ctx context.Context) ([]string, error) {
	return []string{"https://ya.ru", "https://vc.ru", "https://wery-slow-site.com"}, nil
}

type MockAliver struct {
}

func (m *MockAliver) Alive(ctx context.Context, url string) (isAlive bool, responseTime time.Duration) {
	if url == "https://ya.ru" {
		return true, 1000
	}

	if url == "https://vc.ru" {
		return true, 666
	}

	return false, -1
}

func setupCache() *Cache {
	hoster := &MockHoster{}
	aliver := &MockAliver{}
	cache := New(new(sorted.Sort), aliver, hoster, 1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		cache.Watch(ctx)
	}()

	time.Sleep(1 * time.Second)
	cancel()
	return cache
}

type fields struct {
	aliver aliver
	hoster hoster

	mu   *sync.Mutex
	data map[string]dto.Info

	min dto.InfoWithName
	max dto.InfoWithName
	ttl time.Duration
}

func TestCache_GetUrl(t *testing.T) {
	cache := setupCache()

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		want    dto.Info
		wantErr bool
	}{
		{
			name: "Get yandex",
			fields: fields{
				data: cache.data,
			},
			args: args{url: "https://ya.ru"},
			want: dto.Info{
				IsAlive:      true,
				ResponseTime: 1000,
			},
		},
		{
			name: "Get vc",
			fields: fields{
				data: cache.data,
			},
			args: args{url: "https://vc.ru"},
			want: dto.Info{
				IsAlive:      true,
				ResponseTime: 666,
			},
		},
		{
			name: "Get site with 500 error",
			fields: fields{
				data: cache.data,
			},
			args: args{url: "https://wery-slow-site.com"},
			want: dto.Info{
				IsAlive:      false,
				ResponseTime: -1,
			},
		},
		{
			name: "Site not in file",
			fields: fields{
				data: cache.data,
			},
			args:    args{url: "https://any-any-.com"},
			want:    dto.Info{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				ttl:    tt.fields.ttl,
				aliver: tt.fields.aliver,
				hoster: tt.fields.hoster,
				mu:     tt.fields.mu,
				data:   tt.fields.data,
				min:    tt.fields.min,
				max:    tt.fields.max,
			}
			got, err := c.GetURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_GetMax(t *testing.T) {
	cache := setupCache()

	tests := []struct {
		name   string
		want   dto.InfoWithName
		fields fields
	}{
		{
			name: "get max",
			fields: fields{
				data: cache.data,
				max:  cache.max,
			},
			want: dto.InfoWithName{
				Info: dto.Info{
					IsAlive:      true,
					ResponseTime: 1000,
				},
				Name: "https://ya.ru",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				ttl:    tt.fields.ttl,
				aliver: tt.fields.aliver,
				hoster: tt.fields.hoster,
				mu:     tt.fields.mu,
				data:   tt.fields.data,
				min:    tt.fields.min,
				max:    tt.fields.max,
			}
			if got := c.GetMax(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_GetMin(t *testing.T) {
	cache := setupCache()

	tests := []struct {
		name   string
		want   dto.InfoWithName
		fields fields
	}{
		{
			name: "get min",
			fields: fields{
				data: cache.data,
				min:  cache.min,
			},
			want: dto.InfoWithName{
				Info: dto.Info{
					IsAlive:      true,
					ResponseTime: 666,
				},
				Name: "https://vc.ru",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				ttl:    tt.fields.ttl,
				aliver: tt.fields.aliver,
				hoster: tt.fields.hoster,
				mu:     tt.fields.mu,
				data:   tt.fields.data,
				min:    tt.fields.min,
				max:    tt.fields.max,
			}
			if got := c.GetMin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
