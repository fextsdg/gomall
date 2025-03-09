package service

import (
	"context"
	utils2 "gomall/app/frontend/biz/utils"
	"gomall/app/frontend/hertz_gen/frontend/checkout"
	"gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/payment"

	"github.com/cloudwego/hertz/pkg/app"
	rpccheckout "gomall/rpc_gen/kitex_gen/checkout"
)

type CheckOutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckOutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckOutWaitingService {
	return &CheckOutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckOutWaitingService) Run(req *checkout.CheckOutReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc.CheckOutClient.CheckOut(h.Context, &rpccheckout.CheckOutReq{
		UserId:    uint32(utils.GetUserIdFromCtx(h.Context)),
		FirstName: req.GetFirstname(),
		LastName:  req.GetLastname(),
		Email:     req.GetEmail(),
		Address: &rpccheckout.Address{
			Street:  req.GetStreet(),
			City:    req.GetCity(),
			State:   req.GetState(),
			Country: req.GetCountry(),
			ZipCode: req.GetZipcode(),
		},
		CreditInfo: &payment.CreditInfo{
			CreditCardNumber:     req.GetCardNum(),
			CreditCardCvv:        req.GetCvv(),
			CreditExpirationYear: req.GetExpirationYear(),
			CreditExpirationMoth: req.GetExpirationMonth(),
		},
	},
	)
	if err != nil {
		return nil, err
	}
	return utils2.WarpResponse(h.Context, h.RequestContext, map[string]any{"Name": "Waiting", "redirect": "/checkout/result"}), nil
}
