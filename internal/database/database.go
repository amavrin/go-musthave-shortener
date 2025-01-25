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

func NewDB() *DB {
	return &DB{URLs: make(map[string]string)}
}

func makeID(leng int) string {
	once.Do(func() {
		rand.Seed(uint64(time.Now().UnixNano()))
	})
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, leng)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))] // Выбор случайного символа из charset
	}
	return string(result)
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
