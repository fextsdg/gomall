package service

import (
	"context"
	"gomall/app/frontend/infra/rpc"
	"gomall/rpc_gen/kitex_gen/product"

	category "gomall/app/frontend/hertz_gen/frontend/category"

	"github.com/cloudwego/hertz/pkg/app"
)

type ListProductsByCategoryNameService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListProductsByCategoryNameService(Context context.Context, RequestContext *app.RequestContext) *ListProductsByCategoryNameService {
	return &ListProductsByCategoryNameService{RequestContext: RequestContext, Context: Context}
}

func (h *ListProductsByCategoryNameService) Run(req *category.ListProductsReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{CategoryName: req.GetCategory()})
	if err != nil {
		return nil, err
	}

	//// 将[]*Product转换为[]Product
	//productList := make([]product.Product, len(products.GetProducts()))
	//for i, p := range products.GetProducts() {
	//	if p != nil {
	//		productList[i] = *p
	//	}
	//}

	resp = make(map[string]any)
	var ps []product.Product
	for _, v := range products.GetProducts() {
		ps = append(ps, product.Product{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}
	resp["items"] = ps
	resp["Name"] = req.GetCategory()

	return resp, nil

}
