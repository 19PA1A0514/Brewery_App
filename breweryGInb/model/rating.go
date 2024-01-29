package model

import "gorm.io/gorm"

type Rating struct {
	gorm.Model

	Comment string  `json:"comment"`
	Rating  float32 `json:"rating"`
	User    User    `json:"user"`
	UserID  uint    `gorm:"foreignKey:UserID"  json:"userid"`
	Item    Item    `json:"item"`
	ItemID  uint    `gorm:"foreignKey:ItemID" json:"itemid"`
}
