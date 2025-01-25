package database

import (
	"errors"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

var once = sync.Once{}

const randomLeng = 6

type DB struct {
	URLs map[string]string
}

var (
	ErrURLDoesNotExist = errors.New("short URL does not exist")
)

const distinctChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewDB() *DB {
	return &DB{URLs: make(map[string]string)}
}

func encodeOne(b byte) byte {
	return distinctChars[b%byte(len(distinctChars))]
}

func makeID(leng int) string {
	once.Do(func() {
		rand.Seed(uint64(time.Now().UnixNano()))
	})
	buf := make([]byte, leng)
	rand.Read(buf)
	for i := range buf {
		buf[i] = encodeOne(buf[i])
	}
	return string(buf)
}

func (db *DB) SaveURL(URL string) (string, error) {
	for rep := 0; rep < 10; rep++ {
		shortURL := makeID(randomLeng)
		if _, ok := db.URLs[shortURL]; !ok {
			db.URLs[shortURL] = URL
			return shortURL, nil
		}
	}
	return "", errors.New("could not generate short URL")
}

func (db *DB) GetURL(shortURL string) (string, error) {
	URL, ok := db.URLs[shortURL]
	if !ok {
		return "", ErrURLDoesNotExist
	}
	return URL, nil
}
