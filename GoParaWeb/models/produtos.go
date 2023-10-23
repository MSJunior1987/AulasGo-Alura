package models

import (
	"db"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Id, Quantidade  int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	selectDeTodosOsProdutos, erro := db.Query("select * from produtos")

	if erro != nil {
		panic(erro.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		erro = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if erro != nil {
			panic(erro.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close() //executa depois de tudo que executou no metodo conectaComBancoDeDados()
	return produtos
}
