package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/models"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/repositories"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/request"
)

type CategoryService interface {
    CreateCategory(category *request.CategoryRequest) error
    GetCategories() ([]models.Category, error)
    GetCategory(id int) (*models.Category, error)
    UpdateCategory(id int, category *request.CategoryRequest) error
    DeleteCategory(id int) error
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

func (s *CategoryServiceImpl) GetCategories() ([]models.Category, error) {
	return s.CategoryRepository.GetCategories()
}

func (s *CategoryServiceImpl) GetCategory(id int) (*models.Category, error) {
	return s.CategoryRepository.GetCategory(id)
}

func (s *CategoryServiceImpl) UpdateCategory(id int, category *request.CategoryRequest) error {
	return s.CategoryRepository.UpdateCategory(id, category)
}

func (s *CategoryServiceImpl) DeleteCategory(id int) error {
	return s.CategoryRepository.DeleteCategory(id)
}
