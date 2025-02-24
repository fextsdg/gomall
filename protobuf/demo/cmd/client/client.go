package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"gomall/probuf/demo/conf"
	hello "gomall/probuf/demo/kitex_gen/hello"
	"gomall/probuf/demo/kitex_gen/hello/hellosevice"
	"log"
	"time"
)

func main() {
	// 初始化Consul解析器，用于服务发现和客户端负载均衡
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		log.Fatalln(err)
	}
	// 创建Hello服务的客户端，使用Consul进行服务发现，并设置RPC调用的超时时间为3秒
	c := hellosevice.MustNewClient("demoproto", client.WithResolver(r), client.WithRPCTimeout(3*time.Second))
	// 创建上下文对象，用于取消请求、传递请求范围的值等
	cnt := context.Background()
	// 无限循环，每隔一秒向Hello服务发送一个请求
	for {
		resp, err := c.Hello(cnt, &hello.HelloReq{Info: "你好呀！"})
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
