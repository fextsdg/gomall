package service

import (
	"context"
	"gomall/app/frontend/infra/rpc"
	product2 "gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	product "gomall/app/frontend/hertz_gen/frontend/product"
)

type GetProductByIdService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductByIdService(Context context.Context, RequestContext *app.RequestContext) *GetProductByIdService {
	return &GetProductByIdService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductByIdService) Run(req *product.GetProductByIdReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	p, err := rpc.ProductClient.GetProduct(h.Context, &product2.GetProductReq{Id: req.GetId()})
	if err != nil {
		return nil, err
	}
	resp = make(map[string]any)
	resp["item"] = p.Product
	resp["Name"] = "Product"
	return resp, nil
}
