package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/expense-report/expense/database"
	"github.com/nicholasjackson/expense-report/expense/vault"
)

type Expense struct {
	logger hclog.Logger
	db     *database.Client
	vc     *vault.Client
}

func NewExpense(l hclog.Logger, db *database.Client, vc *vault.Client) *Expense {
	return &Expense{l, db, vc}
}

func (e *Expense) HandlePOST(rw http.ResponseWriter, r *http.Request) {
	e.logger.Info("handlers.Expense POST called")

	expense := &database.Expense{}
	err := json.NewDecoder(r.Body).Decode(expense)
	if err != nil {
		e.logger.Error("Unable to process payload", "error", err)
		http.Error(rw, "unable to process payload", http.StatusBadRequest)
		return
	}

	// before inserting encrypt name
	en, err := e.vc.EncryptData(expense.Description.String, "expense_report_service")
	if err != nil {
		e.logger.Error("Unable to encrypt data", "error", err)
		http.Error(rw, "unable to encrypt data", http.StatusBadRequest)
		return
	}

	// set the encrypted name
	expense.Description.String = en

	err = e.db.InsertExpense(expense)
	if err != nil {
		e.logger.Error("Unable to insert expenses", "error", err)
		http.Error(rw, "unable to insert expense items", http.StatusInternalServerError)
		return
	}
}
