package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/masfuulaji/go-expenses-tracker/internal/app/request"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/services"
)

type CategoryHandler interface {
    CreateCategory(w http.ResponseWriter, r *http.Request)
    GetCategories(w http.ResponseWriter, r *http.Request)
    GetCategory(w http.ResponseWriter, r *http.Request)
    UpdateCategory(w http.ResponseWriter, r *http.Request)
    DeleteCategory(w http.ResponseWriter, r *http.Request)
}

type CategoryHandlerImpl struct {
    CategoryService *services.CategoryServiceImpl
}

func NewCategoryHandler(categoryService *services.CategoryServiceImpl) *CategoryHandlerImpl  {
    return &CategoryHandlerImpl{CategoryService: categoryService}
}

func (h *CategoryHandlerImpl) CreateCategory(w http.ResponseWriter, r *http.Request) {
    category := &request.CategoryRequest{}
    err := json.NewDecoder(r.Body).Decode(category)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = h.CategoryService.CreateCategory(category)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "created"})
}

func (h *CategoryHandlerImpl) GetCategories(w http.ResponseWriter, r *http.Request) {
    categories, err := h.CategoryService.GetCategories()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(categories)
}

func (h *CategoryHandlerImpl) GetCategory(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    idn ,err:= strconv.Atoi(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    category, err := h.CategoryService.GetCategory(idn)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandlerImpl) UpdateCategory(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    idn ,err:= strconv.Atoi(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    category := &request.CategoryRequest{}
    err = json.NewDecoder(r.Body).Decode(category)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = h.CategoryService.UpdateCategory(idn, category)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandlerImpl) DeleteCategory(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    idn ,err:= strconv.Atoi(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = h.CategoryService.DeleteCategory(idn)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "deleted"})
}
