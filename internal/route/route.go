package route

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/handlers"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/services"
	"github.com/masfuulaji/go-expenses-tracker/internal/database"
)

func SetupRoutes(r *mux.Router) {
	db, err := database.ConnectDB()
	if err != nil {
        fmt.Println(err)
	}
    if err := db.Ping(context.Background()); err != nil {
        fmt.Println(err)
    }


    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to Ell"))
    }).Methods("GET")


    categoryHandler := handlers.NewCategoryHandler(services.NewCategoryService(db.DB))
    category := r.PathPrefix("/category").Subrouter()
    category.HandleFunc("", categoryHandler.GetCategories).Methods("GET")
    category.HandleFunc("/{id}", categoryHandler.GetCategory).Methods("GET")
    category.HandleFunc("", categoryHandler.CreateCategory).Methods("POST")
    category.HandleFunc("/{id}", categoryHandler.UpdateCategory).Methods("PUT")
    category.HandleFunc("/{id}", categoryHandler.DeleteCategory).Methods("DELETE")
}
