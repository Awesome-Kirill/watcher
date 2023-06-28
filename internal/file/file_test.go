package file

import (
	"context"
	"reflect"
	"testing"
)

func TestFile_Hosts(t *testing.T) {
	type fields struct {
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
				path:  tt.fields.path,
				hosts: tt.fields.hosts,
			}
			got, err := s.Host(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Host() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Host() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Load(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
