package main

import (
	"MAXPUMP1/pkg/di"

	routes "MAXPUMP1/pkg/api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	handler := di.InitializeApi()

	router := gin.Default()

	router = routes.UserRoutes(router, handler)

	router.Run(":8080")

}
