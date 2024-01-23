package main

import (
	"api-go-gin/models"
	"api-go-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Gui Lims", CPF: "34770928807", RG: "404036107"},
		{Nome: "Ana", CPF: "12345678912", RG: "121231233"},
	}
	routes.HandleRequests()
}
