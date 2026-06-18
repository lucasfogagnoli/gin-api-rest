package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

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

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectaBD()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)

	pathDaBusca := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var alunoMock models.Aluno
	json.Unmarshal(response.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaBD()
	CriaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	pathDeDelecao := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", pathDeDelecao, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditaAlunoHandler(t *testing.T) {
	database.ConectaBD()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()

	r.PATCH("/alunos/:id", controllers.EditaAluno)

	aluno := models.Aluno{Nome: "Nome do Aluno Atualizado", CPF: "47345678901", RG: "773456789"}

	alunoJson, _ := json.Marshal(aluno)

	pathDaEdicao := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("PATCH", pathDaEdicao, bytes.NewBuffer(alunoJson))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var alunoMockAtualizado models.Aluno
	json.Unmarshal(response.Body.Bytes(), &alunoMockAtualizado)

	assert.Equal(t, "Nome do Aluno Atualizado", alunoMockAtualizado.Nome)
	assert.Equal(t, "47345678901", alunoMockAtualizado.CPF)
	assert.Equal(t, "773456789", alunoMockAtualizado.RG)
	assert.Equal(t, http.StatusOK, response.Code)
}
