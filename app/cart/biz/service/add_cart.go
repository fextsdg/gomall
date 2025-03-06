package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"gomall/app/cart/biz/dal/mysql"
	"gomall/app/cart/model"
	"gomall/app/cart/rpc"
	cart "gomall/rpc_gen/kitex_gen/cart"
	"gomall/rpc_gen/kitex_gen/product"
)

type AddCartService struct {
	ctx context.Context
} // NewAddCartService new AddCartService
func NewAddCartService(ctx context.Context) *AddCartService {
	return &AddCartService{ctx: ctx}
}

// Run create note info
func (s *AddCartService) Run(req *cart.AddCartReq) (resp *cart.AddCartResp, err error) {
	// Finish your business logic.
	if req.UserId <= 0 {
		return nil, kerrors.NewGRPCBizStatusError(40000, "用户id有误！")
	}
	fmt.Println("111", rpc.ProductClient)
	p, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: int32(req.Item.ProductId)})
	if err != nil {
		return nil, err
	}
	if p == nil || req.Item.ProductId <= 0 {
		return nil, kerrors.NewGRPCBizStatusError(40004, "未找到商品！")
	}
	cq := model.NewCartQuery(mysql.DB, s.ctx)
	err = cq.AddCart(model.Cart{ProductId: req.Item.GetProductId(), UserId: req.GetUserId(), Num: req.Item.GetNum()})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(50001, "添加购物车失败！")
	}
	return
}
