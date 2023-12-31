package produtosController

import (
	"html/template"
	"log"
	"models"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()  /*executa a busca dos produtos*/
	temp.ExecuteTemplate(w, "Index", todosOsProdutos) /*carrega os produtos na pagina index (carregada na pasta templates)*/
}

func Novo(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "novo", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, erro := strconv.ParseFloat(preco, 64)

		if erro != nil {
			log.Println("Erro na conversão do preço", erro)
		}

		quantidadeConvertidaParaInt, erro := strconv.Atoi(quantidade) //converte para inteiro

		if erro != nil {
			log.Println("Erro na conversão da quantidade", erro)
		}
		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertidaParaInt)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("Id")
	models.ExcluiProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("Id")
	produto := models.EditarProduto(idProduto)
	temp.ExecuteTemplate(w, "Editar", produto)
}

func Atualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, erro := strconv.Atoi(id)
		if erro != nil {
			log.Println("Erro na conversão do ID para int:", erro)
		}

		quantidadeConvertidaParaInt, erro := strconv.Atoi(quantidade)
		if erro != nil {
			log.Println("Erro na conversão da quantidade para int:", erro)
		}

		precoConvertidoParaFloat, erro := strconv.ParseFloat(preco, 64)
		if erro != nil {
			log.Println("Erro na conversão da preco para float64:", erro)
		}

		models.AtualizarProduto(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)

	}

	http.Redirect(w, r, "/", 301)
}
