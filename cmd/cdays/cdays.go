package main

import (
	"log"
	"net/http"

	"os"

	"github.com/jumale/cdays/internal/routing"
)

func main() {
	log.Print("The app is starting...")

	r := routing.NewBLRouter()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("The port wasn't set")
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}
