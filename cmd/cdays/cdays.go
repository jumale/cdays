package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("The app is starting...")

	http.HandleFunc("/", rootHandler())

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func rootHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	}
}
