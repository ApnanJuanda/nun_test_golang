package repository

import (
	"api/nun_test/helper"
	"api/nun_test/model"
	"context"
	"database/sql"
)

type OrderRepository interface {
	Save(ctx context.Context, tx *sql.Tx, penjualan *model.Penjualan) (*model.Penjualan, error)
}

type OrderRepositoryImpl struct {
}

func NewOrderRepositoryImpl() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{}
}

func (r OrderRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, penjualan *model.Penjualan) (*model.Penjualan, error) {
	addPenjualanQuery := "INSERT INTO penjualan(nama_pelanggan, tanggal, jam, total, bayar_tunai, kembali) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, addPenjualanQuery, penjualan.NamaPelanggan, penjualan.Tanggal,
		penjualan.Jam, penjualan.Total, penjualan.BayarTunai, penjualan.Kembali)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	penjualanId := int(id)

	if penjualanId > 0 {
		for _, item := range penjualan.Items {
			addPenjualanItemQuery := "INSERT INTO penjualan_item(penjualan_id, item_id, quantity, harga, sub_total) VALUES (?, ?, ?, ?, ?)"
			_, err = tx.ExecContext(ctx, addPenjualanItemQuery, penjualanId, item.ItemId, item.Quantity, item.Harga,
				item.SubTotal)
			helper.PanicIfError(err)
		}
	}
	return penjualan, err
}
