package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("hello")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		//other way below
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("oopss"))
		return
	}
	fmt.Fprintf(rw, "Data %s\n", d)
}
