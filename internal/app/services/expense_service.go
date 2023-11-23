package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/repositories"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/request"
	"github.com/masfuulaji/go-expenses-tracker/internal/app/response"
)

type ExpenseService interface {
	CreateExpense(expense *request.ExpenseRequest) error
	GetExpenses() ([]response.ExpenseResponse, error)
	GetExpense(id string) (*response.ExpenseResponse, error)
	UpdateExpense(id string, expense *request.ExpenseRequest) error
	DeleteExpense(id string) error
}

type ExpenseServiceImpl struct {
	ExpenseRepository *repositories.ExpenseRepositoryImpl
}

func NewExpenseService(db *sqlx.DB) *ExpenseServiceImpl {
	return &ExpenseServiceImpl{ExpenseRepository: repositories.NewExpenseRepository(db)}
}

func (s *ExpenseServiceImpl) CreateExpense(expense *request.ExpenseRequest) error {
	return s.ExpenseRepository.CreateExpense(expense)
}

func (s *ExpenseServiceImpl) GetExpenses() ([]response.ExpenseResponse, error) {
	return s.ExpenseRepository.GetExpenses()
}

func (s *ExpenseServiceImpl) GetExpense(id string) (*response.ExpenseResponse, error) {
	return s.ExpenseRepository.GetExpense(id)
}

func (s *ExpenseServiceImpl) UpdateExpense(id string, expense *request.ExpenseRequest) error {
	return s.ExpenseRepository.UpdateExpense(id, expense)
}

func (s *ExpenseServiceImpl) DeleteExpense(id string) error {
	return s.ExpenseRepository.DeleteExpense(id)
}
