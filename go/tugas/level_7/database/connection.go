package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var tx *sqlx.Tx
var dbdriver string = "sqlite3"
var dbfile string = "../level_5/db/tugas_5.db"

func Connect() *sqlx.DB {
	db := sqlx.MustConnect(dbdriver, dbfile)
	return db
}

func Begin() *sqlx.Tx {
	db, err := sqlx.Open(dbdriver, dbfile)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	tx = db.MustBegin()
	return tx
}

func Rollback() {
	defer tx.Rollback()
}

func Commit() {
	if err := tx.Commit(); err != nil {
		if err = tx.Rollback(); err != nil {
			panic(err)
		}
	}
}
