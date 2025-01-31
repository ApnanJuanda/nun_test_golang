package service

import (
	"api/nun_test/helper"
	"api/nun_test/model"
	"api/nun_test/repository"
	"context"
	"database/sql"
	"strconv"
)

type PenjualanService interface {
	Save(ctx context.Context, request *model.PenjualanRequest) (*model.PenjualanResponse, error)
	GetTotalPriceDetail(ctx context.Context, request *model.TotalPriceRequest) (*model.TotalPriceResponse, error)
	CalculatePriceAfterDiscount(ctx context.Context, request *model.PriceAfterDiscountRequest) (*model.PriceAfterDiscountResponse, error)
}

type PenjualanServiceImpl struct {
	PenjualanRepository repository.PenjualanRepository
	DB                  *sql.DB
}

func NewPenjualanServiceImpl(penjualanRepository repository.PenjualanRepository, DB *sql.DB) *PenjualanServiceImpl {
	return &PenjualanServiceImpl{PenjualanRepository: penjualanRepository, DB: DB}
}

func (s PenjualanServiceImpl) Save(ctx context.Context, request *model.PenjualanRequest) (*model.PenjualanResponse, error) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var penjualanModel = new(model.Penjualan)
	penjualanModel.FromPenjualanRequest(request)
	penjualanModel, err = s.PenjualanRepository.Save(ctx, tx, penjualanModel)
	helper.PanicIfError(err)

	return &model.PenjualanResponse{
		Message: "Success Save Data Penjualan",
	}, nil
}

func (s PenjualanServiceImpl) GetTotalPriceDetail(ctx context.Context, request *model.TotalPriceRequest) (*model.TotalPriceResponse, error) {
	netSales := request.Total / (1 + (request.PersenPajak / 100))
	pajakRp := request.Total - netSales
	return &model.TotalPriceResponse{
		NetSales: netSales,
		PajakRp:  pajakRp,
	}, nil
}

func (s PenjualanServiceImpl) CalculatePriceAfterDiscount(ctx context.Context, request *model.PriceAfterDiscountRequest) (*model.PriceAfterDiscountResponse, error) {
	totalPrice := request.TotalSebelumDiskon
	totalDiskon := 0.0

	for _, d := range request.Discounts {
		discountInt, err := strconv.Atoi(d.Diskon)
		helper.PanicIfError(err)

		diskon := totalPrice * (float64(discountInt) / 100)
		totalDiskon += diskon

		totalPrice -= diskon
	}

	return &model.PriceAfterDiscountResponse{
		TotalDiskon:             totalDiskon,
		TotalHargaSetelahDiskon: totalPrice,
	}, nil
}
