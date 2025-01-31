package service

import (
	"api/nun_test/helper"
	"api/nun_test/model"
	"api/nun_test/repository"
	"context"
	"database/sql"
)

type OrderService interface {
	Save(ctx context.Context, request *model.PenjualanRequest) (*model.PenjualanResponse, error)
	GetTotalPriceDetail(ctx context.Context, request *model.TotalPriceRequest) (*model.TotalPriceResponse, error)
}

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	DB              *sql.DB
}

func NewOrderServiceImpl(orderRepository repository.OrderRepository, DB *sql.DB) *OrderServiceImpl {
	return &OrderServiceImpl{OrderRepository: orderRepository, DB: DB}
}

func (s OrderServiceImpl) Save(ctx context.Context, request *model.PenjualanRequest) (*model.PenjualanResponse, error) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var penjualanModel = new(model.Penjualan)
	penjualanModel.FromPenjualanRequest(request)
	penjualanModel, err = s.OrderRepository.Save(ctx, tx, penjualanModel)
	helper.PanicIfError(err)

	return &model.PenjualanResponse{
		Message: "Success Save Data Penjualan",
	}, nil
}

func (s OrderServiceImpl) GetTotalPriceDetail(ctx context.Context, request *model.TotalPriceRequest) (*model.TotalPriceResponse, error) {
	netSales := request.Total / (1 + (request.PersenPajak / 100))
	pajakRp := request.Total - netSales
	return &model.TotalPriceResponse{
		NetSales: netSales,
		PajakRp:  pajakRp,
	}, nil
}
