package cart

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	cart "gomall/rpc_gen/kitex_gen/cart"
)

func AddCart(ctx context.Context, req *cart.AddCartReq, callOptions ...callopt.Option) (resp *cart.AddCartResp, err error) {
	resp, err = defaultClient.AddCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetCart(ctx context.Context, req *cart.GetCartReq, callOptions ...callopt.Option) (resp *cart.GetCartResp, err error) {
	resp, err = defaultClient.GetCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func EmptyCart(ctx context.Context, req *cart.EmptyCartReq, callOptions ...callopt.Option) (resp *cart.EmptyCartResp, err error) {
	resp, err = defaultClient.EmptyCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "EmptyCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
