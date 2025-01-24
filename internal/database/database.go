package database

import "errors"

type DB struct {
	short2longURLs map[string]string
}

var (
	ErrShortURLAlreadyExists = errors.New("short URL already exists")
	ErrShortURLDoesNotExist  = errors.New("short URL does not exist")
)

func NewDB() *DB {
	return &DB{short2longURLs: make(map[string]string)}
}

func (db *DB) SaveURL(longURL string, shortURL string) error {
	_, ok := db.short2longURLs[shortURL]
	if ok {
		return ErrShortURLAlreadyExists
	}
	db.short2longURLs[shortURL] = longURL
	return nil
}

func (db *DB) GetLongURL(shortURL string) (string, error) {
	longURL, ok := db.short2longURLs[shortURL]
	if !ok {
		return "", ErrShortURLDoesNotExist
	}
	return longURL, nil
}
