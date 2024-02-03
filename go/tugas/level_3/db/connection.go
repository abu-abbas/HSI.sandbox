package db

import (
	"github.com/abu-abbas/level_3/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const driver string = "sqlite3"
const dbfile string = "./db/tugas_3.db"

var tx *sqlx.Tx

func Connect() *sqlx.DB {
	db := sqlx.MustConnect(driver, dbfile)
	return db
}

func Begin() *sqlx.Tx { 
	db, err := sqlx.Open(driver, dbfile)
	utils.ErrorCheck(err)

	err = db.Ping()
	utils.ErrorCheck(err)

	tx = db.MustBegin()

	return tx
}

func Rollback() {
	defer tx.Rollback()
}

func Commit() {
	if err := tx.Commit(); err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			utils.ErrorCheck(errRb)
		}
	}
}
