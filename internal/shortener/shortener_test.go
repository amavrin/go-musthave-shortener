package shortener

import (
	"strings"
	"testing"

	"github.com/amavrin/go-musthave-shortener/internal/database"
	"github.com/stretchr/testify/assert"
)

func TestShortener_shorten(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{
			name: "simple",
			url:  "http://localhost",
		},
		{
			name: "complex",
			url:  "http://localhost:8080/api/v1/shorten",
		},
		{
			name: "empty",
			url:  "",
		},
		{
			name: "long",
			url:  strings.Repeat("a", 100),
		},
		{
			name: "national",
			url:  "http://мойурл.рф",
		},
	}
	for _, tt := range tests {
		results := make(map[string]bool)
		t.Run(tt.name, func(t *testing.T) {
			s := shorten(tt.url)
			assert.Equal(t, len(s), shortURLLength, "shortened url is not the right len")
			_, ok := results[s]
			assert.False(t, ok, "duplicate shortened url")
			results[s] = true
		})
	}
}

func TestShortener_Shorten(t *testing.T) {
	type fields struct {
		db *database.DB
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Shortener{
				db: tt.fields.db,
			}
			got, err := s.Shorten(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shortener.Shorten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Shortener.Shorten() = %v, want %v", got, tt.want)
			}
		})
	}
}
