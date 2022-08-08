package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("hello")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "oops", http.StatusBadRequest)
			//other way below
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("oopss"))
			return
		}
		fmt.Fprintf(rw, "Data %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye")
	})

	// we can use browser localhost:9090
	http.ListenAndServe(":9090", nil)
}
