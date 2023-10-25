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
	selectDeTodosOsProdutos, erro := db.Query("select * from produtos order by id asc")

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

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close() //executa depois de tudo que executou no metodo conectaComBancoDeDados()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, erro := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) value (?, ?, ?, ?)")

	if erro != nil {
		panic(erro.Error())
	}

	/*a variavel e do tipo sql.Stmt*/
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func ExcluiProduto(id string) {
	db := db.ConectaComBancoDeDados()

	querieExcluirProduto, erro := db.Prepare("delete from produtos where id = ?")
	if erro != nil {
		panic(erro.Error())
	}
	querieExcluirProduto.Exec(id)
	defer db.Close()
}

func EditarProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()
	produtoBanco, erro := db.Query("select * from produtos where id = ?", id)
	if erro != nil {
		panic((erro.Error()))
	}

	produtoParaAtualizar := Produto{}

	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		erro = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if erro != nil {
			panic(erro.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Quantidade = quantidade
		produtoParaAtualizar.Preco = preco
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	atualizaProduto, erro := db.Prepare("update produtos set nome=?, descricao=?, preco=?, quantidade=? where id=?")

	if erro != nil {
		panic(erro.Error())
	}

	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
