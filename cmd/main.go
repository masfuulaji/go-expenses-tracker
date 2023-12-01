package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masfuulaji/go-expenses-tracker/internal/route"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	handler := cors.Default().Handler(r)
	route.SetupRoutes(r)

	http.ListenAndServe(":8080", handler)
}
