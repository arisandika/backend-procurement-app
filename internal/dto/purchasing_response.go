package dto

import "time"

type PurchasingResponse struct {
	ID         uint      `json:"id"`
	Date       time.Time `json:"date"`
	GrandTotal float64   `json:"grand_total"`

	Supplier SupplierResponse           `json:"supplier"`
	User     UserResponse               `json:"user"`
	Details  []PurchasingDetailResponse `json:"details"`
}

type SupplierResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type PurchasingDetailResponse struct {
	Item     ItemResponse `json:"item"`
	Qty      int          `json:"qty"`
	Price    float64      `json:"price"`
	SubTotal float64      `json:"sub_total"`
}

type ItemResponse struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
