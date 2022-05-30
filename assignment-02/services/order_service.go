package services

import (
	"all-assignment-scalable-web-service-with-go/assignment-02/models"
	"all-assignment-scalable-web-service-with-go/assignment-02/params"
	"all-assignment-scalable-web-service-with-go/assignment-02/repositories"
	"net/http"
	"time"
)

type OrderService struct {
	orderRepo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) *OrderService {
	return &OrderService{
		orderRepo: repo,
	}
}

func (o *OrderService) CreateOrder(request *params.CreateOrder) *params.Response {
	order := models.Order{
		OrderedAt:    time.Now(),
		CustomerName: request.CustomerName,
		Items:        request.Item,
	}
	err := o.orderRepo.CreateOrder(&order)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request service",
			AdditionalInfo: err.Error(),
		}
	}
	return &params.Response{
		Status:  200,
		Message: "created success",
		Payload: request,
	}

}

func (o *OrderService) GetAllOrders() *params.Response {
	response, err := o.orderRepo.GetAllOrders()
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "internal server error",
			AdditionalInfo: err.Error(),
		}
	}
	return &params.Response{
		Status:  http.StatusOK,
		Payload: response,
	}
}

func (o *OrderService) DeleteOrder(id uint) *params.Response {
	err := o.orderRepo.DeleteOrder(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "order not found",
			AdditionalInfo: err.Error(),
		}
	}
	return &params.Response{
		Status:  http.StatusOK,
		Message: "order success deleted",
	}
}

func (o *OrderService) UpdateOrder(id uint, request *params.CreateOrder) *params.Response {
	model := models.Order{
		OrderedAt:    request.OrderedAt,
		CustomerName: request.CustomerName,
		Items:        request.Item,
	}

	err := o.orderRepo.UpdateOrder(id, &model)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "order not found",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "order success updated",
	}
}
