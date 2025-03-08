package service

import (
	"context"
	"fmt"
	utils2 "github.com/cloudwego/hertz/pkg/common/utils"
	"gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/utils"
	cart2 "gomall/rpc_gen/kitex_gen/cart"
	"gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := utils.GetUserIdFromCtx(h.Context)
	result, err := rpc.CartClient.GetCart(h.Context, &cart2.GetCartReq{UserId: uint32(userId)})
	if result == nil || result.Cart == nil {
		return utils2.H{
			"items":    make([]map[string]any, 0),
			"total":    "0.00",
			"Name":     "Cart",
			"cart_num": 0,
		}, nil
	}
	rows := result.Cart.Items

	var cartItems []map[string]any
	var total float32 = 0.0
	for _, row := range rows {
		p, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: int32(row.ProductId)})
		if err != nil {
			continue
		}
		cartItems = append(cartItems, map[string]any{
			"Name":    p.Product.Name,
			"Picture": p.Product.Picture,
			"Price":   p.Product.Price,
			"Qty":     row.Num,
		})
		total += p.Product.Price * float32(row.Num)
	}

	return utils2.H{
		"items":    cartItems,
		"total":    fmt.Sprintf("%.2f", total),
		"Name":     "Cart",
		"cart_num": len(rows),
	}, nil
}
