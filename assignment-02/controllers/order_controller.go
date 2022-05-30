package controllers

import (
	"all-assignment-scalable-web-service-with-go/assignment-02/params"
	"all-assignment-scalable-web-service-with-go/assignment-02/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{
		orderService: *service,
	}
}

func (o *OrderController) CreateNewOrder(c *gin.Context) {
	req := params.CreateOrder{
		OrderedAt: time.Now(),
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request controller",
			AdditionalInfo: err.Error(),
		})
		return
	}
	response := o.orderService.CreateOrder(&req)
	c.JSON(response.Status, response)
}

func (o *OrderController) GetAllOrders(c *gin.Context) {
	response := o.orderService.GetAllOrders()
	c.JSON(response.Status, response)
}

func (o *OrderController) DeleteOrder(c *gin.Context) {
	id := c.Param("orderId")
	idUint, err1 := strconv.Atoi(id)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request",
			AdditionalInfo: err1.Error(),
		})
		return
	}
	response := o.orderService.DeleteOrder(uint(idUint))
	c.JSON(response.Status, response)

}

func (o *OrderController) UpdateOrder(c *gin.Context) {
	var req *params.CreateOrder
	orderId := c.Param("orderId")

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request",
			AdditionalInfo: err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(orderId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request",
			AdditionalInfo: err.Error(),
		})
		return
	}
	response := o.orderService.UpdateOrder(uint(id), req)
	c.JSON(response.Status, response)
}
