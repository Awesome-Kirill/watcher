package pkg

import "testing"

func TestMakeUrl(t *testing.T) {
	type args struct {
		rawURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test http://", args: args{rawURL: "http://ya.ru"}, want: "http://ya.ru"},
		{name: "test https://", args: args{rawURL: "https://ya.ru"}, want: "https://ya.ru"},
		{name: "test empty", args: args{rawURL: "ya.ru"}, want: "https://ya.ru"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddShema(tt.args.rawURL); got != tt.want {
				t.Errorf("AddShema() = %v, want %v", got, tt.want)
			}
		})
	}
}
