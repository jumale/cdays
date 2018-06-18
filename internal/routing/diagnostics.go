package routing

import (
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

func NewDiagnosticsRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", handleOK())
	r.HandleFunc("/readyz", handleOK())
	return r
}

func handleOK() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusText(http.StatusOK))
	}
}
