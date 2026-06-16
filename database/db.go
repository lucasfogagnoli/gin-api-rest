package database

import (
	"log"

	"github.com/lucasfogagnoli/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBD() {
	stringConexao := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
