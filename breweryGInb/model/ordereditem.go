package model

import "gorm.io/gorm"

type OrderedItem struct {
	gorm.Model

	ItemID   uint `gorm:"foreignkey:ItemID"`
	Quantity int  `json:"quantity"`
	OrderID  uint `gorm:"index" json:"-"`
	Item     Item `json:"item"`
}
