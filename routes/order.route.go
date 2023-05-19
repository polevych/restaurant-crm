package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/polevych/restaurant-crm/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/order", controller.GetOrders())
	incomingRoutes.GET("/order/:order_id", controller.GetOrder())
	incomingRoutes.POST("/order", controller.CreateOrder())
	incomingRoutes.PUT("/order/:order_id", controller.UpdateOrder())
}