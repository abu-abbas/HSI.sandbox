package model

import (
	"database/sql"
	"errors"

	"github.com/abu-abbas/level_5/db"
	"github.com/abu-abbas/level_5/entity"
	"github.com/abu-abbas/level_5/utils"
)

type Detail struct{}

func (d Detail) Migrate() sql.Result {
	var detail entity.Detail

	con := db.Connect()
	res := con.MustExec(detail.Scheme())

	return res
}

func (d Detail) FindById(id int64) (entity.Detail, error) {
	var detail entity.Detail

	qry := "SELECT * FROM item_details WHERE id=?"
	con := db.Connect()
	err := con.Get(&detail, qry, id)

	return d.resultCheck(detail, err)
}

func (d Detail) FindByItemId(item_id int64) ([]entity.Detail, error) {
	var detail []entity.Detail

	qry := "SELECT * FROM item_details WHERE item_id=?"
	con := db.Connect()
	err := con.Select(&detail, qry, item_id)

	if err != nil {
		if err == sql.ErrNoRows {
			return detail, errors.New("item tidak ditemukan")
		} else {
			utils.ErrorCheck(err)
			return detail, err
		}
	}

	return detail, nil
}

func (d Detail) Create(detail entity.Detail) (sql.Result, error) {
	qry := "INSERT INTO item_details (item_id, name) VALUES (:item_id, :name)"
	con := db.Connect()
	res, err := con.NamedExec(qry, detail)

	return res, err
}

func (d Detail) CreateMany(detail []entity.Detail) int64 {
	qry := "INSERT INTO item_details (item_id, name) VALUES (:item_id, :name)"
	trx := db.Begin()
	res, err := trx.NamedExec(qry, detail)
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

func (d Detail) resultCheck(detail entity.Detail, err error) (entity.Detail, error) {
	if err != nil {
		if err == sql.ErrNoRows {
			return detail, errors.New("item tidak ditemukan")
		} else {
			utils.ErrorCheck(err)
			return detail, err
		}
	}

	return detail, nil
}
