package model

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Address struct {
	Street  string
	City    string
	Country string
	State   string
	ZipCode string
}

type Order struct {
	gorm.Model
	UserId       uint32 `gorm:"type:int(11) not null"`
	OrderId      string `gorm:"uniqueIndex;size:256;not null"`
	Email        string
	Address      Address `gorm:"embedded"`
	UserCurrency string
	OrderItems   []*OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

func (o Order) TableName() string {
	return "order"
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) (orders []*Order, err error) {

	err = db.WithContext(ctx).Model(&Order{}).Where("user_id=?", userId).Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
