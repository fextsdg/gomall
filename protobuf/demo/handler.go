package main

import (
	"context"
	"github.com/kitex-contrib/registry-consul/example/hello/kitex_gen/api"
	"gomall/probuf/demo/biz/service"
	hello "gomall/probuf/demo/kitex_gen/hello"
)

// HelloSeviceImpl implements the last service interface defined in the IDL.
type HelloSeviceImpl struct{}

// Hello implements the HelloSeviceImpl interface.
func (s *HelloSeviceImpl) Hello(ctx context.Context, req *hello.HelloReq) (resp *hello.HelloResp, err error) {
	resp, err = service.NewHelloService(ctx).Run(req)

	return resp, err
}
func (h *HelloSeviceImpl) Echo(_ context.Context, req *api.Request) (resp *api.Response, err error) {
	resp = &api.Response{
		Message: req.Message,
	}
	return resp, err
}
