package routes

import (
	"controllers"
	"data"
	"models"
	"net/http"
	"net/http/httptest"
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
