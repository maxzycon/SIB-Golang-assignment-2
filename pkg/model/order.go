package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string
	Item         []Item
}

type Item struct {
	gorm.Model
	ItemCode    string
	Description string
	Quantity    uint64
	OrderID     uint
	Order       Order
}
