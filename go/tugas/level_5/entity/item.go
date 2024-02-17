package entity

import (
	"fmt"
	"net/http"
)

var itemScheme = `
	DROP TABLE IF EXISTS items;

	CREATE TABLE items (
		id		INTEGER			PRIMARY KEY,
		name	VARCHAR(100)	NOT NULL,
		status	VARCHAR(10)		DEFAULT 'draft',
		amount	INTEGER			DEFAULT 0
	);
`

type Item struct {
	Id     int64    `json:"id,omitempty"`
	Name   string   `json:"name"`
	Status string   `json:"status"`
	Amount int      `json:"amount"`
	Detail []Detail `json:"detail,omitempty"`
}

// Bind implements render.Binder.
func (Item) Bind(r *http.Request) error {
	return nil
}

func (item Item) Scheme() string {
	return itemScheme
}

func (item Item) ToString() string {
	return fmt.Sprintf(
		"Item {ID: %d, Name: %s, Status: %s, Amount: %d}",
		item.Id,
		item.Name,
		item.Status,
		item.Amount,
	)
}
