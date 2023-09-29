package main

import (
	"github.com/abdulmanafc2001/testing_ap/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.UserRoutes(router)
	router.Run(":5000")
}
