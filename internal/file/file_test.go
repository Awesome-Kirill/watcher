package file

import (
	"context"
	"reflect"
	"sync"
	"testing"
)

func TestFile_Hosts(t *testing.T) {
	type fields struct {
		once  *sync.Once
		path  string
		hosts []string
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &File{
				once:  tt.fields.once,
				path:  tt.fields.path,
				hosts: tt.fields.hosts,
			}
			got, err := s.Hosts(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hosts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_load(t *testing.T) {
	type fields struct {
		once  *sync.Once
		path  string
		hosts []string
	}

	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &File{
				once:  tt.fields.once,
				path:  tt.fields.path,
				hosts: tt.fields.hosts,
			}
			s.load()
		})
	}
}
