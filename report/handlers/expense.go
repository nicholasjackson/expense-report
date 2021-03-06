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

func (e *Expense) HandleGET(rw http.ResponseWriter, r *http.Request) {
	e.logger.Info("handlers.Expense GET called")

	expenses, err := e.db.GetExpenseItems()
	if err != nil {
		e.logger.Error("Unable to fetch expenses", "error", err)
		http.Error(rw, "unable to fetch expense items", http.StatusInternalServerError)
		return
	}

	// decrypt the data
	decryptedData := []database.Expense{}
	for _, ex := range expenses {
		en, err := e.vc.DecryptData(ex.Description.String, "expense_report_service")
		if err != nil {
			e.logger.Error("Unable to encrypt data", "error", err)
			http.Error(rw, "unable to encrypt data", http.StatusBadRequest)
			return
		}

		// set the decrypted data
		ex.Description.String = en

		// add the decrypted expense to the return collection
		decryptedData = append(decryptedData, ex)
	}

	json.NewEncoder(rw).Encode(decryptedData)
}
