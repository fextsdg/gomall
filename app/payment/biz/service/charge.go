package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	credit "github.com/durango/go-credit-card"
	"github.com/hashicorp/go-uuid"
	"gomall/app/payment/biz/dal/mysql"
	"gomall/app/payment/model"
	payment "gomall/rpc_gen/kitex_gen/payment"
	"strconv"
	"time"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	if req == nil || req.CreditInfo == nil {
		return nil, kerrors.NewGRPCBizStatusError(4004001, "参数错误！")
	}
	creditInfo := req.GetCreditInfo()
	cd := credit.Card{
		Number:  creditInfo.CreditCardNumber,
		Cvv:     strconv.Itoa(int(creditInfo.CreditCardCvv)),
		Month:   strconv.Itoa(int(creditInfo.CreditExpirationMoth)),
		Year:    strconv.Itoa(int(creditInfo.CreditExpirationYear)),
		Company: credit.Company{},
	}
	err = cd.Validate(true) //验证 credit.Card 对象的有效性。true 参数代表测试卡号有效
	if err != nil {
		return nil, err
	}
	transcationId, err := uuid.GenerateUUID()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4004003, err.Error())
	}
	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        uint32(req.UserId),
		OrderId:       req.GetOrderId(),
		TranscationId: transcationId,
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &payment.ChargeResp{TransactionId: transcationId}, nil
}
