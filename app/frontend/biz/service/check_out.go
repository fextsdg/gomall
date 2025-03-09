package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	utils2 "gomall/app/frontend/biz/utils"
	"gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/cart"
	"gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type CheckOutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckOutService(Context context.Context, RequestContext *app.RequestContext) *CheckOutService {
	return &CheckOutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckOutService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	getCart, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{UserId: uint32(utils.GetUserIdFromCtx(h.Context))})
	if err != nil {
		return nil, err
	}
	if getCart == nil || getCart.Cart == nil || getCart.Cart.Items == nil {
		hlog.Error("参数出错！")
	}
	var rows []map[string]any
	var total float32
	items := getCart.Cart.Items
	for _, item := range items {
		getProduct, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: int32(item.GetProductId())})
		if err != nil || getProduct == nil || getProduct.Product == nil {
			continue
		}
		p := getProduct.Product
		rows = append(rows, map[string]any{
			"Picture": p.Picture,
			"Name":    p.Name,
			"Price":   p.Price,
			"Qty":     item.Num,
		})
		total += p.Price * float32(item.Num)
	}

	return utils2.WarpResponse(h.Context, h.RequestContext, map[string]any{
		"total": total,
		"items": rows,
		"Name":  "CheckOut",
	}), nil
}
