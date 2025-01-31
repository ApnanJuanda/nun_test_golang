package model

import (
	"fmt"
	"time"
)

type Penjualan struct {
	ID            int
	NamaPelanggan string
	Tanggal       time.Time
	Jam           time.Time
	Total         float64
	BayarTunai    float64
	Kembali       float64
	Items         []PenjualanItemRequest
}

type PenjualanRequest struct {
	NamaPelanggan string                 `json:"nama_pelanggan"`
	Tanggal       string                 `json:"tanggal"`
	Jam           string                 `json:"jam"`
	Total         float64                `json:"total"`
	BayarTunai    float64                `json:"bayar_tunai"`
	Kembali       float64                `json:"kembali"`
	Items         []PenjualanItemRequest `json:"items"`
}

type PenjualanResponse struct {
	Message string
}

type TotalPriceRequest struct {
	Total       float64 `json:"total"`
	PersenPajak float64 `json:"persen_pajak"`
}

type TotalPriceResponse struct {
	NetSales float64 `json:"net_sales"`
	PajakRp  float64 `json:"pajak_rp"`
	Message  string  `json:"message,omitempty"`
}

type Discount struct {
	Diskon string `json:"diskon"`
}

type PriceAfterDiscountRequest struct {
	Discounts          []Discount `json:"discounts"`
	TotalSebelumDiskon float64    `json:"total_sebelum_diskon"`
}

type PriceAfterDiscountResponse struct {
	TotalDiskon             float64 `json:"total_diskon"`
	TotalHargaSetelahDiskon float64 `json:"total_harga_setelah_diskon"`
	Message                 string  `json:"message,omitempty"`
}

func (m *Penjualan) FromPenjualanRequest(request *PenjualanRequest) {
	tanggalConverted, _ := convertTanggal(request.Tanggal)
	jamConverted, _ := convertJam(request.Tanggal, request.Jam)

	m.NamaPelanggan = request.NamaPelanggan
	m.Tanggal = tanggalConverted
	m.Jam = jamConverted
	m.Total = request.Total
	m.BayarTunai = request.BayarTunai
	m.Kembali = request.Kembali
	m.Items = request.Items
}

func convertTanggal(tanggalStr string) (time.Time, error) {
	layout := "2006-01-02"
	tanggal, err := time.Parse(layout, tanggalStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing tanggal: %v", err)
	}
	return tanggal, nil
}

func convertJam(tanggalStr, jamStr string) (time.Time, error) {
	layout := "2006-01-02 15:04"
	waktu := tanggalStr + " " + jamStr
	jam, err := time.Parse(layout, waktu)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing jam: %v", err)
	}
	return jam, nil
}
