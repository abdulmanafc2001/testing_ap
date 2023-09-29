package routes

import (
	"github.com/abdulmanafc2001/testing_ap/controllers"
	"github.com/abdulmanafc2001/testing_ap/database"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	db := database.InitDatabase()
	router.POST("/", controllers.CreateUser(db))
	router.GET("/:id", controllers.GetUserById(db))
	router.PUT("/:id",controllers.UpdateUser(db))
	router.DELETE("/:id",controllers.DeleteUser(db))
}
