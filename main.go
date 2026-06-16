package main

import (
	"github.com/lucasfogagnoli/gin-api-rest/database"
	"github.com/lucasfogagnoli/gin-api-rest/routes"
)

func main() {
	database.ConectaBD()
	routes.HandleRequests()
}
