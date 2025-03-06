package model

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(255);not null"`
	Description string  `json:"description" gorm:"type:text"`
	Picture     string  `json:"picture" gorm:"type:varchar(255);not null"`
	Price       float32 `json:"price" gorm:"type:decimal(10,2);not null;default:0.0"`
	/**
	many2many:product_category 指定了中间表的名称为 product_category。这个表通常包含两个字段：
	product_id：外键，引用 Product 表的主键。
	category_id：外键，引用 Category 表的主键
	*/
	Categories []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

// 用于查询
type ProductQuery struct {
	db  *gorm.DB
	ctx context.Context
}

func NewProudctQuery(db *gorm.DB, ctx context.Context) ProductQuery {
	return ProductQuery{db: db, ctx: ctx}
}

// 通过商品id查询商品
func (pq ProductQuery) GetProductById(id int32) (p *Product, err error) {
	p = &Product{}
	err = pq.db.WithContext(pq.ctx).First(p, "id=?", id).Error
	return
}

// 通过名称或描述查询商品
func (pq ProductQuery) SearchProducts(query string) (products []*Product, err error) {
	searchPattern := "%" + query + "%"
	err = pq.db.WithContext(pq.ctx).Where("name like ? or description like ?", searchPattern, searchPattern).Find(&products).Error
	return
}
