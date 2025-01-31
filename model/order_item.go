package model

type PenjualanItem struct {
	ID          int
	PenjualanId int
	ItemId      int
	Quantity    float64
	Harga       float64
	SubTotal    float64
}

type PenjualanItemRequest struct {
	ItemId   string  `json:"item_id"`
	Quantity float64 `json:"quantity"`
	Harga    float64 `json:"harga"`
	SubTotal float64 `json:"sub_total"`
}
