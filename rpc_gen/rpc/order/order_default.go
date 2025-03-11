package order

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	order "gomall/rpc_gen/kitex_gen/order"
)

func AddOrder(ctx context.Context, req *order.AddOrderReq, callOptions ...callopt.Option) (resp *order.AddOrderResp, err error) {
	resp, err = defaultClient.AddOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListOrder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (resp *order.ListOrderResp, err error) {
	resp, err = defaultClient.ListOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
