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

	r.Handle("static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Welcome to Ell"))
	}).Methods("GET")

	categoryHandler := handlers.NewCategoryHandler(services.NewCategoryService(db.DB))
	// category := r.PathPrefix("/category").Subrouter()
	r.HandleFunc("/category", categoryHandler.GetCategories).Methods("GET")
	r.HandleFunc("/category/add", categoryHandler.AddCategory).Methods("GET")
	r.HandleFunc("/category/edit/{id}", categoryHandler.EditCategory).Methods("GET")
	r.HandleFunc("/category/{id}", categoryHandler.GetCategory).Methods("GET")
	r.HandleFunc("/category", categoryHandler.CreateCategory).Methods("POST")
	r.HandleFunc("/category/{id}", categoryHandler.UpdateCategory).Methods("PUT")
	r.HandleFunc("/category/{id}", categoryHandler.DeleteCategory).Methods("DELETE")
}
