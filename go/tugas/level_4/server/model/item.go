package model

import (
	"database/sql"
	"errors"

	"github.com/abu-abbas/level_4/server/db"
	"github.com/abu-abbas/level_4/server/entity"
	"github.com/abu-abbas/level_4/server/utils"
)

type Item struct{}

func (i Item) Migrate() sql.Result {
	var item entity.Item

	con := db.Connect()
	res := con.MustExec(item.Scheme())

	return res
}

func (i Item) Get() ([]entity.Item, error) {
	var items []entity.Item

	qry := "SELECT * FROM items"
	con := db.Connect()
	err := con.Select(&items, qry)

	if err != nil {
		if err == sql.ErrNoRows {
			return items, errors.New("item tidak ditemukan")
		} else {
			utils.ErrorCheck(err)
			return items, err
		}
	}

	return items, nil
}

func (i Item) FindById(id int64) (entity.Item, error) {
	var item entity.Item

	qry := "SELECT * FROM items WHERE id=?"
	con := db.Connect()
	err := con.Get(&item, qry, id)

	return i.resultCheck(item, err)
}

func (i Item) Create(item entity.Item) (int64, error) {
	qry := "INSERT INTO items (name, status, amount) VALUES (:name, :status, :amount)"
	con := db.Connect()
	res, err := con.NamedExec(qry, &item)

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (i Item) CreateMany(items []entity.Item) int64 {
	qry := "INSERT INTO items (name, status, amount) VALUES (:name, :status, :amount)"
	trx := db.Begin()
	res, err := trx.NamedExec(qry, items)
	if err != nil {
		trx.Rollback()
	}

	rowAffected, errRowAffected := res.RowsAffected()
	if errRowAffected != nil {
		trx.Rollback()
	}

	trx.Commit()
	return rowAffected
}

func (i Item) CreateWithDetail(item entity.Item) int64 {
	qryItem := "INSERT INTO items (name, status, amount) VALUES (:name, :status, :amount)"
	qryDetail := "INSERT INTO item_details (item_id, name) VALUES (:item_id, :name)"

	trx := db.Begin()
	res, err := trx.NamedExec(qryItem, item)
	if err != nil {
		trx.Rollback()
	}

	lastId, errLastId := res.LastInsertId()
	if errLastId != nil {
		trx.Rollback()
	}

	for idx := range item.Detail {
		item.Detail[idx].ItemId = lastId
	}

	_, err = trx.NamedExec(qryDetail, item.Detail)
	if err != nil {
		trx.Rollback()
	}

	trx.Commit()
	return lastId
}

func (i Item) UpdateItemStatus(item entity.Item) (int64, error) {
	qry := "UPDATE items SET status = :status WHERE id = :id"
	con := db.Connect()
	res, err := con.NamedExec(qry, item)
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}

func (i Item) DeleteItemById(id int64) (int64, error) {
	qry := "DELETE FROM items WHERE id = ?"
	con := db.Connect()
	res := con.MustExec(qry, id)
	return res.RowsAffected()
}

func (i Item) resultCheck(item entity.Item, err error) (entity.Item, error) {
	if err != nil {
		if err == sql.ErrNoRows {
			return item, errors.New("item tidak ditemukan")
		} else {
			utils.ErrorCheck(err)
			return item, err
		}
	}

	return item, nil
}
