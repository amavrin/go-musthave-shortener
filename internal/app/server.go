package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/amavrin/go-musthave-shortener/internal/database"
	"github.com/go-chi/chi/v5"
)

type App struct {
	port    int
	address string
	db      *database.DB
}

const (
	DefaultPort    = 8080
	DefaultAddress = "127.0.0.1"
)

func NewApp(port int, address string) *App {
	db := database.NewDB()
	return &App{port, address, db}
}

func formShortResp(r *http.Request, shortURL string) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	host := r.Host
	return fmt.Sprintf("%s://%s/%s", scheme, host, shortURL)
}

func isValidURL(URL string) bool {
	if len(URL) > 2048 {
		return false
	}
	if !strings.HasPrefix(URL, "http://") && !strings.HasPrefix(URL, "https://") {
		return false
	}
	parsedURL, err := url.ParseRequestURI(URL)
	if err != nil {
		return false
	}
	if parsedURL.Host == "" {
		return false
	}
	return true
}

func (a *App) saveURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	size1MB := int64(1024 * 1024)
	r.Body = http.MaxBytesReader(w, r.Body, size1MB)
	URL, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !isValidURL(string(URL)) {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}
	shortURL, err := a.db.SaveURL(string(URL))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response := formShortResp(r, shortURL)
	w.Write([]byte(response))
}

func (a *App) getURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")
	URL, err := a.db.GetURL(shortURL)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, URL, http.StatusTemporaryRedirect)
}

func (a *App) Run() error {
	r := chi.NewRouter()
	log.Printf("Starting server on %s:%d", a.address, a.port)
	r.Get("/{shortURL}", a.getURL)
	r.Post("/", a.saveURL)
	address := fmt.Sprintf("%s:%d", a.address, a.port)
	return http.ListenAndServe(address, r)
}
