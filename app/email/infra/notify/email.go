package notify

import (
	"github.com/kr/pretty"
	"gomall/rpc_gen/kitex_gen/email"
)

// 模拟一个邮件发送中心
type NoopEmail struct {
}

func (e NoopEmail) Send(req *email.SendReq) {
	pretty.Printf("email:%v\n", req)
}

func NewNoopEmail() *NoopEmail {
	return &NoopEmail{}
}
