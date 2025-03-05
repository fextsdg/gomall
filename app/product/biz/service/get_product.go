package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"gomall/app/product/biz/dal/mysql"
	"gomall/app/product/model"
	product "gomall/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id <= 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "商品ID必须大于0!")
	}
	pq := model.NewProudctQuery(mysql.DB, s.ctx)
	p, err := pq.GetProductById(req.GetId())
	if err != nil {
		return nil, err
	}
	resp = &product.GetProductResp{Product: &product.Product{
		Id:          int32(p.ID),
		Name:        p.Name,
		Description: p.Description,
		Picture:     p.Picture,
		Price:       p.Price,
	}}
	return resp, nil

}
