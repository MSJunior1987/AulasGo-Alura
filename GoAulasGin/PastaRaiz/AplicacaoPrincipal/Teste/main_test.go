// AplicacaoTesteTestes/main_test.go
package AplicacaoTesteTestes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	// Importe o pacote principal aqui

	"github.com/stretchr/testify/assert"
)

func TestRotaPing(t *testing.T) {
	r := main.SetupRotas() // Use o nome do pacote principal para acessar as funções

	// Simula uma solicitação HTTP GET para a rota /ping
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Cria um gravador de resposta falso
	recorder := httptest.NewRecorder()

	// Chama o manipulador da rota
	r.ServeHTTP(recorder, req)

	// Verifica se o código de status da resposta é 200 OK
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Verifica o corpo da resposta
	expectedBody := `{"message":"pong"}`
	assert.Equal(t, expectedBody, recorder.Body.String())
}
