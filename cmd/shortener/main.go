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

func run(port int, address string) error {
	a := app.NewApp(port, address)
	return a.Run()
}

func main() {
	port := flag.Int("port", app.DefaultPort, "port to listen on")
	address := flag.String("address", app.DefaultAddress, "address to listen on")
	flag.Parse()
	err := run(*port, *address)
	if err != nil {
		log.Fatal(err)
	}
}
