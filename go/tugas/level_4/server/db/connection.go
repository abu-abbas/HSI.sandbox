package db

import (
	"github.com/abu-abbas/level_4/server/config"
	"github.com/abu-abbas/level_4/server/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var tx *sqlx.Tx

func Connect() *sqlx.DB {
	db := sqlx.MustConnect(
		config.GetYamlValue().DbConfig.Driver,
		config.GetYamlValue().DbConfig.DbFile,
	)

	return db
}

func Begin() *sqlx.Tx {
	db, err := sqlx.Open(
		config.GetYamlValue().DbConfig.Driver,
		config.GetYamlValue().DbConfig.DbFile,
	)
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
