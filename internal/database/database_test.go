package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB_SaveURL(t *testing.T) {
	type args struct {
		longURL  string
		shortURL string
	}
	tests := []struct {
		name    string
		args    args
		errFunc func(t assert.TestingT, err error, i ...interface{}) bool
	}{
		{
			name: "SaveURL",
			args: args{
				longURL:  "http://example.com",
				shortURL: "abc001",
			},
			errFunc: assert.NoError,
		},
		{
			name: "DuplicatedShort",
			args: args{
				longURL:  "http://example.com",
				shortURL: "abc001",
			},
			errFunc: assert.Error,
		},
		{
			name: "DuplicatedLongShortNotDuplicated",
			args: args{
				longURL:  "http://example.com",
				shortURL: "abc002",
			},
			errFunc: assert.NoError,
		},
	}
	db := NewDB()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := db.SaveURL(tt.args.longURL, tt.args.shortURL)
			tt.errFunc(t, err)
		})
	}
}

func TestDB_GetLongURL(t *testing.T) {
	type args struct {
		longURL  string
		shortURL string
	}
	tests := []struct {
		name    string
		args    args
		errFunc func(t assert.TestingT, err error, i ...interface{}) bool
	}{
		{
			name: "URL",
			args: args{
				longURL:  "http://example.com",
				shortURL: "abc001",
			},
			errFunc: assert.NoError,
		},
	}
	db := NewDB()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := db.SaveURL(tt.args.longURL, tt.args.shortURL)
			tt.errFunc(t, err)
			got, err := db.GetLongURL(tt.args.shortURL)
			tt.errFunc(t, err)
			assert.Equal(t, tt.args.longURL, got, "GetLongURL() = %v, want %v", got, tt.args.longURL)
		})
	}
}
