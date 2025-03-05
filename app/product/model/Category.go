package model

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string    `json:"categoryName" gorm:"type:varchar(40) not null"`
	Description  string    `json:"description" gorm:"type:text"`
	Product      []Product `json:"product" gorm:"many2many:product_category"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	db  *gorm.DB
	ctx context.Context
}

// 通过分类名称获取商品列表
func (cq CategoryQuery) GetProductByCategoryName(categoryName string) (categorites []Category, err error) {
	err = cq.db.WithContext(cq.ctx).Where(&Category{CategoryName: categoryName}).Preload("product").Find(&categorites).Error
	return
}
