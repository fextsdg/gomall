package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ProductId uint32 `gorm:"type :int(11) not null"`
	UserId    uint32 `gorm:"type :int(11) not null;index:idx_user_id"`
	Num       uint32 `gorm:"type :int(11) not null"`
}

func (c Cart) TableName() string {
	return "cart"
}

type CartQuery struct {
	db  *gorm.DB
	ctx context.Context
}

func NewCartQuery(db *gorm.DB, ctx context.Context) CartQuery {
	return CartQuery{db: db, ctx: ctx}
}

func (cq CartQuery) AddCart(item Cart) (err error) {
	//先检查购物车中是否已存在该商品，若存在更新数量即可
	result := Cart{}
	err = cq.db.WithContext(cq.ctx).Model(&Cart{}).Where(&Cart{
		ProductId: item.ProductId,
		UserId:    item.UserId,
	}).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if result.ID > 0 {
		return cq.db.WithContext(cq.ctx).Model(&Cart{}).Where(&Cart{
			ProductId: item.ProductId,
			UserId:    item.UserId,
		}).UpdateColumn("num", gorm.Expr("num+?", item.Num)).Error
	}

	return cq.db.WithContext(cq.ctx).Model(&Cart{}).Create(&item).Error
}

func (cq CartQuery) GetCart(userId uint32) (items []*Cart, err error) {
	err = cq.db.WithContext(cq.ctx).Model(&Cart{}).Where("user_id=?", userId).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (cq CartQuery) EmptyCart(userId uint32) (err error) {
	if userId <= 0 {
		return errors.New("用户id有误！")
	}
	return cq.db.WithContext(cq.ctx).Delete(&Cart{}, "user_id=?", userId).Error //软删除
	//return cq.db.WithContext(cq.ctx).Unscoped().Delete(&Cart{}, "user_id=?", userId).Error //硬删除
}
