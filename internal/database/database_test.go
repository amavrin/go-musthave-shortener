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

func TestDB_GetURL(t *testing.T) {
	tests := []struct {
		name    string
		URL     string
		short   string
		errFunc func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool
	}{
		{
			name:    "get non existing url",
			short:   "non-existing-url",
			errFunc: assert.Error,
		},
		{
			name:    "get existing url",
			URL:     "existing-url",
			errFunc: assert.NoError,
		},
	}
	db := NewDB()
	for _, tt := range tests {
		if tt.URL != "" {
			short, err := db.SaveURL(tt.URL)
			assert.NoError(t, err, "SaveURL() error = %v", err)
			tt.short = short
		}
		t.Run(tt.name, func(t *testing.T) {
			_, err := db.GetURL(tt.short)
			tt.errFunc(t, err, "GetURL() error = %v", err)
		})
	}
}
