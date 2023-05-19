package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/polevych/restaurant-crm/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/order-item", controller.GetOrderItems())
	incomingRoutes.GET("/order-item/order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.GET("/order-item/:order-item_id", controller.GetOrderItem())
	incomingRoutes.POST("/order-item", controller.CreateOrderItem())
	incomingRoutes.PUT("/order-item/:order-item_id", controller.UpdateOrderItem())
}