package service

import (
	"context"
	hello "gomall/probuf/demo/kitex_gen/hello"
	"testing"
)

func TestHello_Run(t *testing.T) {
	ctx := context.Background()
	s := NewHelloService(ctx)
	// init req and assert value

	req := &hello.HelloReq{Info: "你好！"}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
