package dto

import "time"

type PaloadOrder struct {
	OrderAt      time.Time           `json:"orderedAt"`
	CustomerName string              `json:"customerName"`
	Items        []PayloadItemsOrder `json:"items"`
}

type PayloadItemsOrder struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Qty         uint64 `json:"quantity"`
}

type OrderResponse struct {
	ID                uint                `json:"id"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
	CustomerName      string              `json:"customer_name"`
	OrderItemResponse []OrderItemResponse `json:"items"`
}

type OrderItemResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ItemCode    string    `json:"itemcode"`
	Description string    `json:"description"`
	Qty         uint64    `json:"quantity"`
	OrderID     uint      `json:"orderid"`
}
