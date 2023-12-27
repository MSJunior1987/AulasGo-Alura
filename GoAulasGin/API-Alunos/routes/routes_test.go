package routes

import (
	"bytes"
	"controllers"
	"data"
	"encoding/json"
	"fmt"
	"models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) /*simplifica as mensagens de teste*/
	rotas := gin.Default()
	return rotas
}

var ID int

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome Aluno Teste", CPF: "12345678901", RG: "123456787"}
	data.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	data.DB.Delete(&aluno, ID)
}

func TestVerificarStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	requisicao, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder() /*interface que guarda informacoes da requisicao*/
	r.ServeHTTP(resposta, requisicao)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
}

func TestVerificarStatusCodeDaSaudacaoComParametroRetornaJSON(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	requisicao, _ := http.NewRequest("GET", "/marcio", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, requisicao)

	if !assert.Equal(t, http.StatusOK, resposta.Code) {
		t.Errorf("Os StatusCodes não são iguais. Esperado: %d, Recebido: %d", http.StatusOK, resposta.Code)
	}

	mockDaResposta := `{"API diz":"E ai marcio, tudo beleza?"}`
	respostaBody := resposta.Body
	if !assert.Equal(t, mockDaResposta, string(respostaBody.Bytes())) {
		t.Errorf("Os corpos das respostas não são iguais. Esperado: %s, Recebido: %s", mockDaResposta, string(respostaBody.Bytes()))
	}
}

func TestListandoTodosAlunosHandler(t *testing.T) {
	data.ConectaComBancoDeDados()
	// CriaAlunoMock()
	// defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	// fmt.Println(resposta.Body)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	data.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	data.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	caminhoDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", caminhoDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	fmt.Println(alunoMock.Nome)
	assert.Equal(t, "Nome Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456787", alunoMock.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletaAlunoHandler(t *testing.T) {
	data.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	caminhoDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", caminhoDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaAlunoHandler(t *testing.T) {
	data.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "Nome Aluno Teste", CPF: "01987654321", RG: "123456789"} //struct para edicao
	valorJason, _ := json.Marshal(aluno)                                                 //converte a struc em json
	caminhoParaEditar := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", caminhoParaEditar, bytes.NewBuffer(valorJason))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "01987654321", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456789", alunoMockAtualizado.RG)
	assert.Equal(t, "Nome Aluno Teste", alunoMockAtualizado.Nome)
}
