package main

import (
	"data"
	"routes"
)

func main() {
	data.ConectaComBancoDeDados()
	routes.HandleRequests()
}
