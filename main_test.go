package main

import (
	"fmt"
	"io"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucasfogagnoli/gin-api-rest/controllers"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/Lucas", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	/* Método antes de utilizar o 'testify'
	if response.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", response.Code, http.StatusOK)
	}*/

	/* Método utilizando o 'testify' */
	assert.Equal(t, http.StatusOK, response.Code, "Deveriam ser iguais")

	mockDaResponse := `{"API diz":"E ai Lucas, tudo beleza?"}`
	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, mockDaResponse, string(responseBody))
	fmt.Println(string(responseBody))
	fmt.Println(mockDaResponse)
}
