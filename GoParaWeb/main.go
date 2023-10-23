package main

import (
	"net/http"
	"routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil) //cria o servidor go na porta 8000 (http://locahost:8000/)
}
