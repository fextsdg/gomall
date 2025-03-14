package service

import (
	"context"
	"errors"
	"gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/types"
	"gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/order"
	product2 "gomall/rpc_gen/kitex_gen/product"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	listOrder, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{UserId: uint32(utils.GetUserIdFromCtx(h.Context))})
	if err != nil {
		return nil, err
	}
	if listOrder == nil {
		return nil, errors.New("订单列表为空！")
	}
	var orders []types.Order
	orderItems := listOrder.GetOrders()
	for _, item := range orderItems {
		var oi []types.Item
		for _, i2 := range item.OrderItems {
			productResp, err := rpc.ProductClient.GetProduct(h.Context, &product2.GetProductReq{Id: int32(i2.CartItem.ProductId)})
			if err != nil || productResp == nil || productResp.GetProduct() == nil {
				continue
			}
			oi = append(oi, types.Item{
				Picture:     productResp.GetProduct().Picture,
				ProductName: productResp.GetProduct().GetName(),
				Qty:         i2.CartItem.GetNum(),
				Cost:        i2.GetCost(),
			})
		}
		orders = append(orders, types.Order{
			CreatedDate: time.Unix(item.CreatedAt, 0).Format("2006-01-02 15:04:05"),
			OrderId:     item.OrderId,
			Items:       oi,
		})
	}
	return map[string]any{
		"orders": orders,
		"Name":   "Order",
	}, nil
}
