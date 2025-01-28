package app

import "testing"

func Test_isValidURL(t *testing.T) {
	type args struct {
		URL string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid URL",
			args: args{URL: "http://example.com"},
			want: true,
		},
		{
			name: "invalid URL",
			args: args{URL: "******************"},
			want: false,
		},
		{
			name: "empty URL",
			args: args{URL: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidURL(tt.args.URL); got != tt.want {
				t.Errorf("isValidURL(%s) = %v, want %v", tt.args.URL, got, tt.want)
			}
		})
	}
}
