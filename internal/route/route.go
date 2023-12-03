package route

import (
	"context"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/handlers"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/services"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/utils"
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

	r.HandleFunc("/dashboard", func(w http.ResponseWriter, _ *http.Request) {
		tmpl, err := template.ParseFiles("views/dashboard.html")
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}).Methods("GET")

	categoryHandler := handlers.NewCategoryHandler(services.NewCategoryService(db.DB))
	category := r.PathPrefix("/category").Subrouter()
	category.HandleFunc("", categoryHandler.GetCategories).Methods("GET")
	category.HandleFunc("/add", categoryHandler.AddCategory).Methods("GET")
	category.HandleFunc("/edit/{id}", categoryHandler.EditCategory).Methods("GET")
	category.HandleFunc("/{id}", categoryHandler.GetCategory).Methods("GET")
	category.HandleFunc("", categoryHandler.CreateCategory).Methods("POST")
	category.HandleFunc("/{id}", categoryHandler.UpdateCategory).Methods("PUT")
	category.HandleFunc("/{id}", categoryHandler.DeleteCategory).Methods("DELETE")
}
