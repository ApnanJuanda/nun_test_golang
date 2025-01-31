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
	GetTotalPriceDetail(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CalculatePriceAfterDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
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

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	response, err := c.OrderService.Save(request.Context(), &penjualanRequest)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		response.Message = "Terjadi Internal server error"
		err = encoder.Encode(response)
		helper.PanicIfError(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = encoder.Encode(response)
	helper.PanicIfError(err)
}

func (c OrderControllerImpl) GetTotalPriceDetail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var priceRequest = model.TotalPriceRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&priceRequest)
	helper.PanicIfError(err)

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	response, err := c.OrderService.GetTotalPriceDetail(request.Context(), &priceRequest)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		response.Message = "Terjadi Internal server error"
		err = encoder.Encode(response)
		helper.PanicIfError(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = encoder.Encode(response)
	helper.PanicIfError(err)
}

func (c OrderControllerImpl) CalculatePriceAfterDiscount(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var priceDiscountRequest = model.PriceAfterDiscountRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&priceDiscountRequest)
	helper.PanicIfError(err)

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	response, err := c.OrderService.CalculatePriceAfterDiscount(request.Context(), &priceDiscountRequest)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		response.Message = "Terjadi Internal server error"
		err = encoder.Encode(response)
		helper.PanicIfError(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	err = encoder.Encode(response)
	helper.PanicIfError(err)
}
