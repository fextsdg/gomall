package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/cart"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

// 用于添加用户登录信息
func WarpResponse(ctx context.Context, c *app.RequestContext, resp map[string]any) map[string]any {
	// todo edit custom code
	// 确保 resp 不为 nil
	if resp == nil {
		resp = make(map[string]any)
	}
	userId := utils.GetUserIdFromCtx(ctx)
	resp["user_id"] = userId

	cartResp, err := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{UserId: uint32(userId)})
	if err != nil || cartResp == nil || cartResp.Cart == nil {
		return resp
	}
	resp["cart_num"] = len(cartResp.Cart.Items)
	return resp
}
