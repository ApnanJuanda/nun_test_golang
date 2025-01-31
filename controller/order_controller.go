package controller

import (
	"api/nun_test/helper"
	"api/nun_test/model"
	"api/nun_test/service"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type OrderController interface {
	Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderControllerImpl(orderService service.OrderService) *OrderControllerImpl {
	return &OrderControllerImpl{OrderService: orderService}
}

func (c OrderControllerImpl) Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var penjualanRequest = model.PenjualanRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&penjualanRequest)
	helper.PanicIfError(err)

	response, err := c.OrderService.Save(request.Context(), &penjualanRequest)
	if err != nil {
		response.Message = "Terjadi Internal server error"
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(response)
	helper.PanicIfError(err)
}
