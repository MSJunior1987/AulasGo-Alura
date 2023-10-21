package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conectaComBancoDeDados() *sql.DB {
	// conexao := ""
	db, erro := sql.Open("mysql", "user:root@tcp(localhost:3306)/alura_loja")
	if erro != nil {
		panic(erro.Error())
	}
	return db
}

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	db := conectaComBancoDeDados()
	defer db.Close() //executa depois de tudo que executou no metodo conectaComBancoDeDados()
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil) //acessa com o endereco http://localhost/8000 no modo anonimo do chrome
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortavel", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto Novo", "Muito Legal", 1.99, 1},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
