package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() *sql.DB {
	db, erro := sql.Open("mysql", "root:root@tcp(localhost:3306)/db_bons_d06")
	if erro != nil {
		panic(erro.Error())
	}
	return db
}
