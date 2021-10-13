package database

type Expense struct {
	ID           int     `json:"id" db:"id"`
	Name         string  `json:"name" db:"name"`
	TripID       string  `json:"tripId" db:"trip_id"`
	Cost         float32 `json:"cost" db:"cost"`
	Currency     string  `json:"currency" db:"currency"`
	Date         string  `json:"date" db:"date"`
	Reimbursable bool    `json:"reimbursable" db:"reimbursable"`
}
