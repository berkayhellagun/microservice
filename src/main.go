package main

import (
	"log"
	"net/http"
	"os"

	"github.com/berkayhellagun/microservice/src/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	// serve mux is multiplexer
	mux := http.NewServeMux()
	mux.Handle("/", hh)

	// we can use browser localhost:9090
	// first parameter bind address and second parameter is the handler
	// if we do not have any http handler system automaticly going default serve mux
	// but we have ServeMux configuration we can use it as second parameter
	http.ListenAndServe(":9090", mux)
}
