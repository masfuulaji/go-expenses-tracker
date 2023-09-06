package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masfuulaji/go-expenses-tracker/internal/route"
)

func main() {
	r := mux.NewRouter()

	route.SetupRoutes(r)

	http.ListenAndServe(":8080", r)
}
