//go:build tests

package sorted

import (
	"reflect"
	"testing"
	"watcher/internal/dto"
)

func TestSort_MinMax(t *testing.T) {
	d := make(map[string]dto.Info)

	d["a.ru"] = dto.Info{
		IsAlive:      true,
		ResponseTime: 1,
	}

	d["b.ru"] = dto.Info{
		IsAlive:      true,
		ResponseTime: 15,
	}

	d["c.ru"] = dto.Info{
		IsAlive:      true,
		ResponseTime: 14,
	}

	d["d.ru"] = dto.Info{
		IsAlive:      false,
		ResponseTime: 28,
	}

	tests := []struct {
		name    string
		c       *Sort
		data    map[string]dto.Info
		wantMin dto.InfoWithName
		wantMax dto.InfoWithName
	}{
		{
			name: "",
			c:    new(Sort),
			data: d,
			wantMin: dto.InfoWithName{
				Name: "a.ru",
				Info: dto.Info{
					IsAlive:      true,
					ResponseTime: 1,
				},
			},

			wantMax: dto.InfoWithName{
				Name: "b.ru",
				Info: dto.Info{
					IsAlive:      true,
					ResponseTime: 15,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := tt.c.MinMax(tt.data)
			if !reflect.DeepEqual(gotMin, tt.wantMin) {
				t.Errorf("MinMax() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if !reflect.DeepEqual(gotMax, tt.wantMax) {
				t.Errorf("MinMax() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
