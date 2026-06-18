package main

import (
	"fmt"
	"io"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucasfogagnoli/gin-api-rest/controllers"
	"github.com/lucasfogagnoli/gin-api-rest/database"
	"github.com/lucasfogagnoli/gin-api-rest/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
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

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaBD()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaBD()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}
