package main

import (
	"context"
	"gomall/app/email/biz/service"
	
)

// EmailServiceImpl implements the last service interface defined in the IDL.
type EmailServiceImpl struct{}

// Send implements the EmailServiceImpl interface.
func (s *EmailServiceImpl) Send(ctx context.Context, req *email.SendReq) (resp *email.SendResp, err error) {
	resp, err = service.NewSendService(ctx).Run(req)

	return resp, err
}
