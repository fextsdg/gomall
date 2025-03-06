package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"gomall/app/cart/biz/dal/mysql"
	"gomall/app/cart/model"
	cart "gomall/rpc_gen/kitex_gen/cart"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Finish your business logic.
	if req.UserId <= 0 {
		return nil, kerrors.NewGRPCBizStatusError(40000, "用户ID有误！")
	}
	cq := model.NewCartQuery(mysql.DB, s.ctx)
	err = cq.EmptyCart(req.UserId)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(50002, "清空购物车失败！")
	}
	return
}
