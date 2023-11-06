package main

import (
	"data"
	"fmt"
	"models"
	"routes"
)

func main() {
	//acessa via o link 'http://localhost:8000/api/personalidades'
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	data.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
