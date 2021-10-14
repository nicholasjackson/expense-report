package database

import (
	"gopkg.in/guregu/null.v3"
)

type Expense struct {
	ID           []uint8     `json:"id" db:"id"`
	Name         string      `json:"name" db:"name"`
	Description  null.String `json:"description" db:"description"`
	TripID       string      `json:"tripId" db:"trip_id"`
	Cost         float32     `json:"cost" db:"cost"`
	Currency     string      `json:"currency" db:"currency"`
	Date         string      `json:"date" db:"date"`
	Reimbursable bool        `json:"reimbursable" db:"reimbursable"`
}
