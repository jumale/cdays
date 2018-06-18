package main

import "net/http"

func main() {
	http.HandleFunc("/", rootHandler())
	http.ListenAndServe(":8000", nil)
}

func rootHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
