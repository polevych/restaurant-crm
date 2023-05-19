package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/polevych/restaurant-crm/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/table", controller.GetTables())
	incomingRoutes.GET("/table/:table_id", controller.GetTable())
	incomingRoutes.POST("/table", controller.CreateTable())
	incomingRoutes.PUT("/table/:table_id", controller.UpdateTable())
}