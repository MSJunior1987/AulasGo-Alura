package routes

import (
	produtosController "controllers"
	"net/http"
)

func CarregaRotas() {
	/*chama uma pagina html definida pelo endereco*/
	http.HandleFunc("/", produtosController.Index)
	http.HandleFunc("/novo", produtosController.Novo)
	http.HandleFunc("/insert", produtosController.Insert)
	http.HandleFunc("/excluir", produtosController.Delete)
}
