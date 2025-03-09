package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type PaymentLog struct {
	gorm.Model
	UserId        uint32    `json:"user_id"`
	OrderId       string    `json:"order_id"`
	TranscationId string    `json:"transcation_id"`
	Amount        float32   `json:"amount" gorm:"type:DECIMAL(10,2)"`
	PayAt         time.Time `json:"pay_at"`
}

func (pl PaymentLog) TableName() string {
	return "payment_log"
}

func CreatePaymentLog(db *gorm.DB, ctx context.Context, log *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(log).Error
}
