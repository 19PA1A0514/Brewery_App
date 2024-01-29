package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderedTime time.Time     `json:"orderedTime"`
	Status      OrderStatus   `json:"status"`
	User        User          `json:"user" gorm:"foreignKey:UserID"`
	UserID      uint          `validate:"required"`
	Latitude    string        `validate:"required" json:"latitude"`
	Longitude   string        `validate:"required" json:"longitude"`
	Invoice     Invoice       `json:"-"`
	InvoiceID   uint          `gorm:"foreignKey:InvoiceID" json:"-"`
	FoodItems   []OrderedItem `gorm:"foreignKey:OrderID" validate:"required" json:"fooditems"`
	TotalAmount float32
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "Pending"
	OrderStatusConfirmed OrderStatus = "Confirmed"
	OrderStatusShipped   OrderStatus = "Shipped"
	OrderStatusDelivered OrderStatus = "Delivered"
)
