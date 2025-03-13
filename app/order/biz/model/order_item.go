package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductId    uint32
	OrderIdRefer string `gorm:"size:256;index;not null"`
	Quantity     int32
	Cost         float32
}

func (i OrderItem) TableName() string {
	return "order_item"
}
