package email

import (
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"gomall/app/email/infra/mq"
	"gomall/app/email/infra/notify"
	"gomall/rpc_gen/kitex_gen/email"
	"google.golang.org/protobuf/proto"
)

// 初始化消费者
func ConsumerInit() {
	//订阅主题
	subscribe, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.SendReq
		err := proto.Unmarshal(msg.Data, &req) //反序列化
		if err != nil {
			return
		}
		noopEamil := notify.NewNoopEmail()
		noopEamil.Send(&req) //发送邮件
	})
	if err != nil {
		panic(err)
	}
	//注册一个服务关闭时的钩子函数，用于取消订阅和关闭连接
	server.RegisterShutdownHook(func() {
		subscribe.Unsubscribe()
		mq.Nc.Close()
	})
}
