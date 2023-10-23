package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaComBancoDeDados() *sql.DB {
	// conexao := ""
	db, erro := sql.Open("mysql", "root:root@tcp(localhost:3306)/alura_loja")
	if erro != nil {
		panic(erro.Error())
	}
	return db
}
