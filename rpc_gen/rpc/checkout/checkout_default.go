package checkout

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	checkout "gomall/rpc_gen/kitex_gen/checkout"
)

func CheckOut(ctx context.Context, req *checkout.CheckOutReq, callOptions ...callopt.Option) (resp *checkout.CheckOutResp, err error) {
	resp, err = defaultClient.CheckOut(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CheckOut call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
