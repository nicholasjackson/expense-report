package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/expense-report/expense/database"
)

type Expense struct {
	logger hclog.Logger
	db     *database.Client
}

func NewExpense(l hclog.Logger, db *database.Client) *Expense {
	return &Expense{l, db}
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

	err = e.db.InsertExpense(expense)
	if err != nil {
		e.logger.Error("Unable to insert expenses", "error", err)
		http.Error(rw, "unable to insert expense items", http.StatusInternalServerError)
		return
	}
}
