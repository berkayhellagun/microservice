package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/berkayhellagun/microservice/src/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//handlers
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	// serve mux is multiplexer
	mux := http.NewServeMux()
	mux.Handle("/", hh)
	mux.Handle("/goodbye", gh)
	mux.Handle("/products", ph)
	// we can use browser localhost:9090
	// first parameter bind address and second parameter is the handler
	// if we do not have any http handler system automaticly going default serve mux
	// but we have ServeMux configuration we can use it as second parameter

	// create custom server with configuration
	s := &http.Server{
		Addr:         ":9090",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//http.ListenAndServe(":9090", mux)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Shutdown", sig)
	// no longer accept any more requests but its gonna wait until finished the work
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
