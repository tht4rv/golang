package main

import (
	// Import the gorilla/mux library we just installed

	"INSTALLAMS/handlers"
	"net/http"
	"github.com/gorilla/mux"
	"INSTALLAMS/middleware/logmiddleware"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/statusams",logmiddleware.LogMiddleware(http.HandlerFunc(handlers.StatusAms))).Methods("POST")
	r.Handle("/installams",logmiddleware.LogMiddleware(http.HandlerFunc(handlers.InstallAMS))).Methods("POST")
	r.Handle("/uninstallams",logmiddleware.LogMiddleware(http.HandlerFunc(handlers.UninstallAMS))).Methods("POST")
	r.Handle("/selectams",logmiddleware.LogMiddleware(http.HandlerFunc(handlers.SelectAMS))).Methods("POST")
	http.ListenAndServe(":8080", r)
}
