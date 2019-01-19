package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"rsc.io/quote"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, quote.Hello())
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.Handle("/", r)
	fmt.Printf("Listening")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
