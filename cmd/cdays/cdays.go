package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Print("The app is starting...")

	r := mux.NewRouter()
	r.HandleFunc("/home", rootHandler())

	log.Fatal(http.ListenAndServe(":8000", r))
}

func rootHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	}
}
