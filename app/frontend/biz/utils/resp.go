package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"gomall/app/frontend/utils"
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
	resp["user_id"] = utils.GetUserIdFromCtx(ctx)
	return resp
}
