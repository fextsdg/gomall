package service

import (
	"context"
	"gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/hertz/pkg/app"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type CheckOutResultService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckOutResultService(Context context.Context, RequestContext *app.RequestContext) *CheckOutResultService {
	return &CheckOutResultService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckOutResultService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc.CartClient.EmptyCart(h.Context, &cart.EmptyCartReq{UserId: uint32(utils.GetUserIdFromCtx(h.Context))})
	if err != nil {
		return nil, err
	}
	return map[string]any{"Name": "Result"}, nil
}
