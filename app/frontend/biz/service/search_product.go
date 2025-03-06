package service

import (
	"context"
	"gomall/app/frontend/infra/rpc"
	product2 "gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	product "gomall/app/frontend/hertz_gen/frontend/product"
)

type SearchProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductService(Context context.Context, RequestContext *app.RequestContext) *SearchProductService {
	return &SearchProductService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	products, err := rpc.ProductClient.SearchProducts(h.Context, &product2.SearchProductsReq{Query: req.GetQuery()})
	if err != nil {
		return nil, err
	}
	resp = make(map[string]any)
	resp["items"] = products.Products
	resp["Name"] = "Search"
	resp["q"] = req.GetQuery()

	return
}
