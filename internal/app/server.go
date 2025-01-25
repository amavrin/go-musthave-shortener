package app

import (
	"fmt"
	"io"
	"log"
	"net/http"

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

func (a *App) saveURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	URL, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shortURL, err := a.db.SaveURL(string(URL))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))
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
	log.Print("here**************")

	r := chi.NewRouter()
	r.Post("/", a.saveURL)
	r.Get("/{shortURL}", a.getURL)
	address := fmt.Sprintf("%s:%d", a.address, a.port)
	http.ListenAndServe(address, r)
	return nil
}
