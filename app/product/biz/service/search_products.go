package service

import (
	"context"
	"gomall/app/product/biz/dal/mysql"
	"gomall/app/product/model"
	product "gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProudctQuery(mysql.DB, s.ctx)
	result, err := productQuery.SearchProducts(req.GetQuery())
	if err != nil {
		return nil, err
	}
	var products []*product.Product
	for _, v := range result {
		products = append(products, &product.Product{
			Id:          int32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}
	resp = &product.SearchProductsResp{Products: products}
	return resp, nil
}
