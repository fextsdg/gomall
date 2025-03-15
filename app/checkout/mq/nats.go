package mq

import "github.com/nats-io/nats.go"

//初始化nats中间件

var (
	err error
	Nc  *nats.Conn
)

func Init() {
	Nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
}
