package database

import (
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	db *sqlx.DB
}

func New(connection string) (*Client, error) {
	db, err := sqlx.Connect("mysql", connection)
	if err != nil {
		return nil, err
	}

	return &Client{db}, nil
}

func (c *Client) GetExpenseItems() ([]Expense, error) {
	expenses := []Expense{}

	err := c.db.Select(&expenses, "SELECT * FROM expense_item;")
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (c *Client) InsertExpense(expense *Expense) error {
	uuid := uuid.New()
	id := strings.Replace(uuid.String(), "-", "", -1)

	sql := `
	INSERT INTO expense_item (id, name, trip_id, cost, currency, date, reimbursable)
	VALUES (?,?,?,?,?,?,?)`

	_, err := c.db.Exec(sql, id[:16], expense.Name, expense.TripID, expense.Cost, expense.Currency, expense.Date, expense.Reimbursable)

	return err
}
