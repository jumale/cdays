package main

import (
	"log"
	"net/http"

	"github.com/jumale/cdays/internal/routing"
)

func main() {
	log.Print("The app is starting...")

	r := routing.NewBLRouter()

	log.Fatal(http.ListenAndServe(":8000", r))
}

func rootHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	}
}
