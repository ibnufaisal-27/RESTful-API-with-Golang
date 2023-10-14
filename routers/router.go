package routers

import (
	"github.com/gin-gonic/gin"
	"mekari/controllers"
)

func SetupRouter() *gin.Engine {
	route := gin.Default()
	route.POST("/employees", controllers.Create)
	route.GET("/employees", controllers.GetEmployees)
	route.GET("/employees/:id", controllers.FindEmployee)
	route.PUT("/employees/:id", controllers.UpdateEmployee)
	route.DELETE("/employees/:id", controllers.DeleteEmployee)
	return route

}
