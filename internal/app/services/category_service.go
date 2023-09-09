package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/repositories"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/request"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/response"
)

type CategoryService interface {
    CreateCategory(category *request.CategoryRequest) error
    GetCategories() ([]response.CategoryResponse, error)
    GetCategory(id string) (*response.CategoryResponse, error)
    UpdateCategory(id string, category *request.CategoryRequest) error
    DeleteCategory(id string) error
}
type CategoryServiceImpl struct {
    CategoryRepository *repositories.CategoryRepositoryImpl
}

func NewCategoryService(db *sqlx.DB) *CategoryServiceImpl {
    return &CategoryServiceImpl{CategoryRepository: repositories.NewCategoryRepository(db)}
}

func (s *CategoryServiceImpl) CreateCategory(category *request.CategoryRequest) error {
	return s.CategoryRepository.CreateCategory(category)
}

func (s *CategoryServiceImpl) GetCategories() ([]response.CategoryResponse, error) {
	return s.CategoryRepository.GetCategories()
}

func (s *CategoryServiceImpl) GetCategory(id string) (*response.CategoryResponse, error) {
	return s.CategoryRepository.GetCategory(id)
}

func (s *CategoryServiceImpl) UpdateCategory(id string, category *request.CategoryRequest) error {
	return s.CategoryRepository.UpdateCategory(id, category)
}

func (s *CategoryServiceImpl) DeleteCategory(id string) error {
	return s.CategoryRepository.DeleteCategory(id)
}
