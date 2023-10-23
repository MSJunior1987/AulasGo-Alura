package routes

import (
	produtosController "controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", produtosController.Index)
}
