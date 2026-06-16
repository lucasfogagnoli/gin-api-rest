package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasfogagnoli/gin-api-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)

	r.Run()
}
