package model

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model

	Amount float32 `json:"amount"`

	DeliveryFee float32 `json:"delivery-fee"`

	Gst float32 `json:"gst"`

	TotalAmount float32 `json:"total-amount"`
}
