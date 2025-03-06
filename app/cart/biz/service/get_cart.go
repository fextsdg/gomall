package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"gomall/app/cart/biz/dal/mysql"
	"gomall/app/cart/model"
	cart "gomall/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {

	// Finish your business logic.
	if req.GetUserId() <= 0 {
		return nil, kerrors.NewGRPCBizStatusError(40000, "用户Id有误！")
	}
	cq := model.NewCartQuery(mysql.DB, s.ctx)
	items, err := cq.GetCart(req.GetUserId())
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(50002, "查询购物车清单失败！")
	}
	var rows []*cart.CartItem
	for _, item := range items {
		rows = append(rows, &cart.CartItem{
			ProductId: item.ProductId,
			Num:       item.Num,
		})
	}

	return &cart.GetCartResp{Cart: &cart.Cart{
		UserId: req.GetUserId(),
		Items:  rows,
	}}, nil
}
