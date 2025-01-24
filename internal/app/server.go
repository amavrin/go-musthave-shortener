package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct {
	port    int
	address string
}

const (
	DefaultPort    = 8080
	DefaultAddress = "127.0.0.1"
)

func NewApp(port int, address string) *App {
	return &App{port, address}
}

func (a *App) Run() error {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	address := fmt.Sprintf("%s:%d", a.address, a.port)
	http.ListenAndServe(address, r)
	return nil
}
