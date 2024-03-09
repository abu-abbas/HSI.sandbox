package database

import (
	"database/sql"
	"errors"
)

type Item struct{}

func (i Item) Get() ([]ItemEntity, error) {
	var items []ItemEntity

	query := "SELECT * FROM items"
	con := Connect()
	err := con.Select(&items, query)

	if err != nil {
		if err == sql.ErrNoRows {
			return items, errors.New("item tidak ditemukan")
		} else {
			return items, err
		}
	}

	return items, nil
}

func (i Item) UpdateAmount(item ItemEntity) (int64, error) {
	query := "UPDATE items SET amount = :amount WHERE amount != :amount"
	con := Connect()
	res, err := con.NamedExec(query, item)
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}
