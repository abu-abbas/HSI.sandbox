package entity

import "fmt"

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
	Id		int64			`db:"id"`
	Name	string			`db:"name"`
	Status	string			`db:"status"`
	Amount	int				`db:"amount"`
	Detail	[]Detail
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
