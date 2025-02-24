package service

import (
	"context"
	hello "gomall/probuf/demo/kitex_gen/hello"
)

type HelloService struct {
	ctx context.Context
} // NewHelloService new HelloService
func NewHelloService(ctx context.Context) *HelloService {
	return &HelloService{ctx: ctx}
}

// Run create note info
func (s *HelloService) Run(req *hello.HelloReq) (resp *hello.HelloResp, err error) {
	// Finish your business logic.

	return &hello.HelloResp{Info: req.Info}, nil
}
