package main

import (
	"api/nun_test/config"
	"api/nun_test/controller"
	"api/nun_test/helper"
	"api/nun_test/repository"
	"api/nun_test/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db := config.GetMyConnection()
	penjualanRepository := repository.NewPenjualanRepositoryImpl()
	penjualanService := service.NewPenjualanServiceImpl(penjualanRepository, db)
	penjualanController := controller.NewPenjualanControllerImpl(penjualanService)

	router := httprouter.New()
	router.POST("/api/penjualan", penjualanController.Save)
	router.GET("/api/penjualan/price-detail", penjualanController.GetTotalPriceDetail)
	router.POST("/api/penjualan/price-discount", penjualanController.CalculatePriceAfterDiscount)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	fmt.Println("My Application is running")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
