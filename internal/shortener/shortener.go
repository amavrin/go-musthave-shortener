package shortener

import (
	"crypto/sha256"

	"github.com/amavrin/go-musthave-shortener/internal/database"
)

const (
	shortURLLength  = 6
	shortenAttempts = 3
)

type Shortener struct {
	db *database.DB
}

func NewShortener(db *database.DB) *Shortener {
	return &Shortener{db}
}

func shorten(url string) string {
	hash := sha256.Sum256([]byte(url))
	shortURL := string(hash[:shortURLLength])
	return shortURL
}

func (s *Shortener) Shorten(url string) (string, error) {
	shortURL := shorten(url)
	err := s.db.SaveURL(shortURL, url)
	for i := 0; i < shortenAttempts && err != nil; i++ {
		shortURL = shorten(shortURL)
		err = s.db.SaveURL(shortURL, url)
	}
	return shortURL, err
}

func (s *Shortener) Expand(shortURL string) (string, error) {
	return s.db.GetLongURL(shortURL)
}
