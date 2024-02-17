package entity

import "fmt"

var detailScheme = `
	DROP TABLE IF EXISTS item_details;

	CREATE TABLE item_details (
		id		INTEGER			PRIMARY KEY,
		item_id	INTEGER			NOT NULL,
		name	VARCHAR(10)		NOT NULL
	);
`

type Detail struct {
    Id      int64   `db:"id"`
    ItemId  int64   `db:"item_id"`
    Name    string  `db:"name"`
}

func (detail Detail) ToString() string {
	return fmt.Sprintf("Detail {ID: %d, Item ID: %d, Detail Name: %s}\n", detail.Id, detail.ItemId, detail.Name)
}

func (detail Detail) Scheme() string {
	return detailScheme
}

