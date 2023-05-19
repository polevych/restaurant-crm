package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/polevych/restaurant-crm/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/menu", controller.GetMenus())
	incomingRoutes.GET("/menu/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menu", controller.CreateMenu())
	incomingRoutes.PUT("/menu/:menu_id", controller.UpdateMenu())
}