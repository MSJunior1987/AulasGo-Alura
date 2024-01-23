package routes

import (
	"api-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorCodigo)
	r.DELETE("/alunos/:id", controllers.ExcluiAluno)
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	r.Run()
}
