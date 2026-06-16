package main

import (
	"github.com/lucasfogagnoli/gin-api-rest/database"
	"github.com/lucasfogagnoli/gin-api-rest/models"
	"github.com/lucasfogagnoli/gin-api-rest/routes"
)

func main() {
	database.ConectaBD()
	models.Alunos = []models.Aluno{
		{Nome: "Lucas Tavares", CPF: "012.345.678-90", RG: "01.234.567-8"},
		{Nome: "Vanessa Silver", CPF: "012.345.678-91", RG: "01.234.567-9"},
	}
	routes.HandleRequests()
}
