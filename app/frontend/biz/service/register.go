package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	"github.com/cloudwego/hertz/pkg/app"
	auth "gomall/app/frontend/hertz_gen/frontend/auth"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 123)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
