package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/models"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/request"
)

var category models.Category

type CategoryRepository interface {
	CreateCategory(category *request.CategoryRequest) error
	GetCategories() ([]models.Category, error)
	GetCategory(id int) (*models.Category, error)
	UpdateCategory(id int, category *request.CategoryRequest) error
	DeleteCategory(id int) error
}

type CategoryRepositoryImpl struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepositoryImpl {
    return &CategoryRepositoryImpl{db: db}
}

func (r *CategoryRepositoryImpl) CreateCategory(category *request.CategoryRequest) error {
	query := `INSERT INTO category (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4)`
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	err := r.db.QueryRow(query, category.Name, category.Description, createdAt, updatedAt).Scan()
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepositoryImpl) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	query := `SELECT * FROM category WHERE deleted_at IS NULL`

	err := r.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepositoryImpl) GetCategory(id int) (*models.Category, error) {
	var category models.Category
	query := `SELECT * FROM category WHERE id = $1 AND deleted_at IS NULL`
	err := r.db.Get(&category, query, id)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) UpdateCategory(id int, category *request.CategoryRequest) error {
	query := `UPDATE category SET name = $1, description = $2, updated_at = $3 WHERE id = $4`
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := r.db.Exec(query, category.Name, category.Description, updatedAt, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepositoryImpl) DeleteCategory(id int) error {
	query := `UPDATE category SET deleted_at = $1 WHERE id = $2`
	deletedAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := r.db.Exec(query, deletedAt, id)
	if err != nil {
		return err
	}
	return nil
}
