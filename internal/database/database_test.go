package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB_SaveURL(t *testing.T) {

	tests := []struct {
		name string
		URL  string
	}{
		{
			name: "SaveURL",
			URL:  "http://example.com",
		},
		{
			name: "SaveURL2",
			URL:  "http://example2.com",
		},
		{
			name: "SaveURL3",
			URL:  "http://example3.com",
		},
	}
	db := NewDB()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			short, err := db.SaveURL(tt.URL)
			assert.NoError(t, err, "SaveURL() error = %v", err)
			assert.Equal(t, len(short), randomLeng, "shortened url is not the right len")
			original, err := db.GetURL(short)
			assert.NoError(t, err, "GetURL() error = %v", err)
			assert.Equal(t, original, tt.URL, "GetURL() = %v, want %v", original, tt.URL)
		})
	}
}
