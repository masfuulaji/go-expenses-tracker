package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/models"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/request"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/response"
)

var expense models.Expense

type ExpenseRepository interface {
	CreateExpense(expense *request.ExpenseRequest) error
	GetExpenses() ([]response.ExpenseResponse, error)
	GetExpense(id string) (*response.ExpenseResponse, error)
	UpdateExpense(id string, expense *request.ExpenseRequest) error
	DeleteExpense(id string) error
}

type ExpenseRepositoryImpl struct {
	db *sqlx.DB
}

func NewExpenseRepository(db *sqlx.DB) *ExpenseRepositoryImpl {
	return &ExpenseRepositoryImpl{db: db}
}

func (r *ExpenseRepositoryImpl) CreateExpense(expense *request.ExpenseRequest) error {
	query := `INSERT INTO expense (date, description, category_id, incoming, outgoing, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := r.db.Exec(query, expense.Date, expense.Description, expense.CategoryId, expense.Incoming, expense.Outgoing, expense.Balance, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *ExpenseRepositoryImpl) GetExpenses() ([]response.ExpenseResponse, error) {
	expenses := []response.ExpenseResponse{}
	query := `SELECT id, date, description, category_id, incoming, outgoing, balance, created_at, updated_at FROM expense WHERE deleted_at IS NULL`
	err := r.db.Select(&expenses, query)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func (r *ExpenseRepositoryImpl) GetExpense(id string) (*response.ExpenseResponse, error) {
	var expense response.ExpenseResponse
	query := `SELECT id, date, description, category_id, incoming, outgoing, balance, created_at, updated_at FROM expense WHERE id = $1 AND deleted_at IS NULL`
	err := r.db.Get(&expense, query, id)
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func (r *ExpenseRepositoryImpl) UpdateExpense(id string, expense *request.ExpenseRequest) error {
	query := `UPDATE expense SET date = $1, description = $2, category_id = $3, incoming = $4, outgoing = $5, balance = $6, updated_at = $7 WHERE id = $8`
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := r.db.Exec(query, expense.Date, expense.Description, expense.CategoryId, expense.Incoming, expense.Outgoing, expense.Balance, updatedAt, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ExpenseRepositoryImpl) DeleteExpense(id string) error {
	query := `UPDATE expense SET deleted_at = $1 WHERE id = $2`
	deletedAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := r.db.Exec(query, deletedAt, id)
	if err != nil {
		return err
	}
	return nil
}
