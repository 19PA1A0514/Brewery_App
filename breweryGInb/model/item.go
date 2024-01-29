package model

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name            string   `validate:"required" json:"name"`
	Details         string   `validate:"required" json:"details"`
	Price           float32  `validate:"required" json:"price"`
	InStock         bool     `json:"-"`
	Rating          float32  `json:"-"`
	QuantityOrdered int      `json:"-"`
	Type            FoodType `validate:"required" json:"foodtype"`
	Image           string   `validate:"required" json:"image"`
}

type FoodType string

const (
	FoodTypeCraft    FoodType = "Craft"
	FoodTypeDraught  FoodType = "Draught"
	FoodTypeImported FoodType = "Imported"
	FoodTypeSnacks   FoodType = "Snacks"
)
