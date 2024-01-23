package main

import (
	"api-go-gin/data"
	"api-go-gin/routes"
)

func main() {
	data.ConectaComBancoDeDados()
	routes.HandleRequests()
}
