package sorted

import (
	"reflect"
	"testing"
	"watcher/internal/dto"
)

func TestSort_MinMax(t *testing.T) {
	type args struct {
		data map[string]dto.Info
	}
	tests := []struct {
		name    string
		c       Sort
		args    args
		wantMin dto.InfoWithName
		wantMax dto.InfoWithName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := tt.c.MinMax(tt.args.data)
			if !reflect.DeepEqual(gotMin, tt.wantMin) {
				t.Errorf("MinMax() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if !reflect.DeepEqual(gotMax, tt.wantMax) {
				t.Errorf("MinMax() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
