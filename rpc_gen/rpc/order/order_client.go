package order

import (
	"context"
	order "gomall/rpc_gen/kitex_gen/order"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	
)

type RPCClient interface {
	KitexClient() orderservice.Client
	Service() string
	AddOrder(ctx context.Context, Req *order.AddOrderReq, callOptions ...callopt.Option) (r *order.AddOrderResp, err error)
	ListOrder(ctx context.Context, Req *order.ListOrderReq, callOptions ...callopt.Option) (r *order.ListOrderResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := orderservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient orderservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() orderservice.Client {
	return c.kitexClient
}

func (c *clientImpl) AddOrder(ctx context.Context, Req *order.AddOrderReq, callOptions ...callopt.Option) (r *order.AddOrderResp, err error) {
	return c.kitexClient.AddOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) ListOrder(ctx context.Context, Req *order.ListOrderReq, callOptions ...callopt.Option) (r *order.ListOrderResp, err error) {
	return c.kitexClient.ListOrder(ctx, Req, callOptions...)
}
