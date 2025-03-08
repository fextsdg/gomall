package service

import (
	"context"
	"fmt"
	"gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/utils"
	cart2 "gomall/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/hertz/pkg/app"
	cart "gomall/app/frontend/hertz_gen/frontend/cart"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := utils.GetUserIdFromCtx(h.Context)

	_, err = rpc.CartClient.AddCart(h.Context, &cart2.AddCartReq{
		UserId: uint32(userId),
		Item: &cart2.CartItem{
			ProductId: req.ProductId,
			Num:       req.Num,
		},
	})
	fmt.Println(userId, req)
	return
}
