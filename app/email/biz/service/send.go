package service

import (
	"context"
	email "gomall/rpc_gen/kitex_gen/email"
)

type SendService struct {
	ctx context.Context
} // NewSendService new SendService
func NewSendService(ctx context.Context) *SendService {
	return &SendService{ctx: ctx}
}

// Run create note info
func (s *SendService) Run(req *email.SendReq) (resp *email.SendResp, err error) {
	// Finish your business logic.

	return
}
