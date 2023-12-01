package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/request"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/services"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/utils"
)

type CategoryHandler interface {
	CreateCategory(w http.ResponseWriter, r *http.Request)
	GetCategories(w http.ResponseWriter, r *http.Request)
	GetCategory(w http.ResponseWriter, r *http.Request)
	EditCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
}

type CategoryHandlerImpl struct {
	CategoryService *services.CategoryServiceImpl
}

func NewCategoryHandler(categoryService *services.CategoryServiceImpl) *CategoryHandlerImpl {
	return &CategoryHandlerImpl{CategoryService: categoryService}
}

func (h *CategoryHandlerImpl) CreateCategory(w http.ResponseWriter, r *http.Request) {
	category := &request.CategoryRequest{}
	err := json.NewDecoder(r.Body).Decode(category)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.CategoryService.CreateCategory(category)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "created"})
}

func (h *CategoryHandlerImpl) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.CategoryService.GetCategories()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tmpl, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = tmpl.Execute(w, categories)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *CategoryHandlerImpl) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	category, err := h.CategoryService.GetCategory(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, category)
}

func (h *CategoryHandlerImpl) AddCategory(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/category/form.html")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *CategoryHandlerImpl) EditCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	category, err := h.CategoryService.GetCategory(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tmpl, err := template.ParseFiles("views/category/form.html")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = tmpl.Execute(w, category)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *CategoryHandlerImpl) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	category := &request.CategoryRequest{}
	err := json.NewDecoder(r.Body).Decode(category)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.CategoryService.UpdateCategory(id, category)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "updated"})
}

func (h *CategoryHandlerImpl) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.CategoryService.DeleteCategory(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "deleted"})
}
