package main

import "github.com/gin-gonic/gin"

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Marcio Junior",
	})
}

func main() {
	r := gin.Default() //cria um servidor com as configuracoes padrao
	r.GET("/alunos", ExibeTodosAlunos)

	r.Run() //inicializa o servidor
	//r.Run("5000") //inicializa o servidor na porta 5000
}
