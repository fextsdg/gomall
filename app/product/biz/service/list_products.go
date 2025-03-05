package service

import (
	"context"
	"gomall/app/product/biz/dal/mysql"
	"gomall/app/product/model"
	product "gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	cateQuery := model.NewCategoryQuery(mysql.DB, s.ctx)
	result, err := cateQuery.GetProductByCategoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}
	var products []*product.Product
	for _, v := range result {
		for _, v1 := range v.Product {
			products = append(products, &product.Product{
				Id:          int32(v1.ID),
				Name:        v1.Name,
				Description: v1.Description,
				Picture:     v1.Picture,
				Price:       v1.Price,
			})
		}
	}
	return &product.ListProductsResp{Products: products}, nil
}
