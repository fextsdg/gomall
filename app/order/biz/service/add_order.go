package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hashicorp/go-uuid"
	"gomall/app/order/biz/dal/mysql"
	"gomall/app/order/biz/model"
	order "gomall/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type AddOrderService struct {
	ctx context.Context
} // NewAddOrderService new AddOrderService
func NewAddOrderService(ctx context.Context) *AddOrderService {
	return &AddOrderService{ctx: ctx}
}

// Run create note info
func (s *AddOrderService) Run(req *order.AddOrderReq) (resp *order.AddOrderResp, err error) {
	// Finish your business logic.
	if len(req.OrderItems) == 0 {
		return nil, kerrors.NewGRPCBizStatusError(6004001, "订单列表不存在！")
	}
	db := mysql.DB
	orderId, _ := uuid.GenerateUUID()
	//一个事务：
	//先创建订单-->再创建订单信息列表
	err = db.Transaction(func(tx *gorm.DB) error {

		reqAddress := req.GetAddress()
		var address model.Address
		if reqAddress != nil {
			address = model.Address{
				Street:  reqAddress.GetStreet(),
				City:    reqAddress.GetCity(),
				Country: reqAddress.GetCountry(),
				State:   reqAddress.GetState(),
				ZipCode: reqAddress.GetZipCode(),
			}
		}
		err = tx.Model(&model.Order{}).Create(&model.Order{

			UserId:       req.UserId,
			OrderId:      orderId,
			Address:      address,
			UserCurrency: req.GetUserCurrency(),
			Email:        req.Email,
		}).Error
		if err != nil {
			return err
		}
		var itemList []model.OrderItem
		//创建订单列表信息
		for _, item := range req.OrderItems {
			if item.CartItem == nil {
				continue
			}
			itemList = append(itemList, model.OrderItem{
				ProductId:    item.CartItem.ProductId,
				OrderIdRefer: orderId,
				Quantity:     int32(item.CartItem.Num),
				Cost:         item.Cost,
			})
		}
		err = tx.Model(&model.OrderItem{}).Create(itemList).Error
		if err != nil {
			return err
		}
		return nil
	})
	return &order.AddOrderResp{AddOrderResult: &order.AddOrderResult{OrderId: orderId}}, nil
}
