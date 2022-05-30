package main

import (
	"all-assignment-scalable-web-service-with-go/assignment-02/controllers"
	"all-assignment-scalable-web-service-with-go/assignment-02/database"
	"all-assignment-scalable-web-service-with-go/assignment-02/repositories"
	"all-assignment-scalable-web-service-with-go/assignment-02/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDB()
	router := gin.Default()

	orderRepo := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	router.POST("/orders", orderController.CreateNewOrder)
	router.GET("/orders", orderController.GetAllOrders)
	router.PUT("/orders/:orderId", orderController.UpdateOrder)
	router.DELETE("/orders/:orderId", orderController.DeleteOrder)
	router.Run(database.APP_PORT)
}
