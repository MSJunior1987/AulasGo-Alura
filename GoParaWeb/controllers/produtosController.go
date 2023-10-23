package produtosController

import (
	"html/template"
	"models"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()  /*executa a busca dos produtos*/
	temp.ExecuteTemplate(w, "Index", todosOsProdutos) /*carrega os produtos na pagina index (carregada na pasta templates)*/
}
