package main

import (
	"flag"
	"log"

	"github.com/amavrin/go-musthave-shortener/internal/app"
)

type Config struct {
	Port    int
	Address string
}

func run(c Config) error {
	a := app.NewApp(c.Port, c.Address)
	return a.Run()
}

func main() {
	config := Config{}
	config.Port = *flag.Int("port", app.DefaultPort, "port to listen on")
	config.Address = *flag.String("address", app.DefaultAddress, "address to listen on")
	flag.Parse()
	err := run(config)
	if err != nil {
		log.Fatal("failed to run server: %w", err)
	}
}
