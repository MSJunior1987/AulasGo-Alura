package controllers

import (
	"api-go-gin/data"
	"api-go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	data.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{"API diz": "E ai " + nome + ", tudo beleza?"})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	//ShouldBindJSON - pega o corpo de entrada de acordou com a struct de aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}
	data.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCodigo(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	data.DB.First(&aluno, id)
	if ValidaIDRetorno(int(aluno.ID)) {
		c.JSON(http.StatusOK, aluno)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"Not found": "Aluno não encontrado"})
}

func ExcluiAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	data.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{"data": "Aluno excluído com sucesso!"})
}

func AtualizaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	data.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	data.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	data.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if ValidaIDRetorno(int(aluno.ID)) {
		c.JSON(http.StatusOK, aluno)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"Not found": "Aluno não encontrado"})
}

func ValidaIDRetorno(id int) bool {
	if id != 0 {
		return true
	}
	return false
}
