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

	cl := &Client{db}
	cl.performMigrations()

	return cl, nil
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
	INSERT INTO expense_item (id, name, description, trip_id, cost, currency, date, reimbursable)
	VALUES (?,?,?,?,?,?,?,?)`

	_, err := c.db.Exec(
		sql,
		id[:16],
		expense.Name,
		expense.Description,
		expense.TripID,
		expense.Cost,
		expense.Currency,
		expense.Date,
		expense.Reimbursable,
	)

	return err
}

func (c *Client) performMigrations() {
	sql := `
ALTER TABLE expense_item
ADD COLUMN description TEXT AFTER name;
`
	c.db.Exec(sql)
}
