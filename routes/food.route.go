package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/polevych/restaurant-crm/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/food", controller.GetFoods())
	incomingRoutes.GET("/food/:food_id", controller.GetFood())
	incomingRoutes.POST("/food", controller.CreateFood())
	incomingRoutes.PUT("/food/:food_id", controller.UpdateFood())
}